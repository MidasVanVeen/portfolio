.PHONY: build templ tailwindcss

run: build tailwindcss
	./a.out

build: templ
	go build -o a.out

templ-watch:
	templ generate --watch

templ:
	templ generate

tailwindcss:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

tailwindcss-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify --watch
