package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/midasvanveen/portfolio/v2/db"
	"github.com/midasvanveen/portfolio/v2/handlers"
	"github.com/midasvanveen/portfolio/v2/middleware"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port         string `envconfig:"PORT" default:":4000"`
	DatabaseName string `envconfig:"DATABASE" default:"portfolio.db"`
}

func loadConfig() *Config {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := chi.NewRouter()

	cfg := loadConfig()

	database := db.MustOpen(cfg.DatabaseName)

	resumeStore := db.NewResumeStore(database)
	projectsStore := db.NewProjectStore(database)
	// projectsStore.CreateProject("UNLOCK Biocontroller", "The UNLOCK biocontroller project was initiated to address the lack of flexibility in commercially available biocontrollers used in bioreactor cultivation. Traditional biocontrollers fall short in meeting the specific needs of researchers pushing the boundaries of cultivation techniques and processes.", "https://gitlab.com/m-unlock/pcp/bio-c-kernel")
	// resumeStore.CreateResumeEntry("UNLOCK Biocontroller", "https://m-unlock.nl/", "2023-present", "Software Engineer", []db.ResumeLink{
	// 	{Title: "Lead the development of the GUI"},
	// 	{Title: "Wrote drivers for industrial sensors and actuators"},
	// })
	//
	{
		staticServer := http.FileServer(http.Dir("./static"))
		r.Handle("/static/*", http.StripPrefix("/static/", staticServer))
	}

	r.Group(func(r chi.Router) {
		r.Use(
			chi_middleware.Logger,
			middleware.TextHTMLMiddleware,
			middleware.CSPMiddleware,
		)

		r.NotFound(handlers.NotFoundHandler)
		r.HandleFunc("/", handlers.IndexHandler)
		r.HandleFunc("/resume", handlers.NewResumeHandler(resumeStore).ServeHTTP)
		r.HandleFunc("/gallery", handlers.NewGalleryHandler(projectsStore).ServeHTTP)
		r.HandleFunc("/contact", handlers.ContactHandler)
	})

	killSig := make(chan os.Signal, 1)

	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			logger.Info("Server shutdown complete")
		} else if err != nil {
			logger.Error("Server error", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	logger.Info("Server started", slog.String("port", cfg.Port))

	<-killSig

	logger.Info("Shutting down server")

	// Create a context with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}
}
