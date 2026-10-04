FROM docker.io/library/golang:1.22-bookworm AS build

WORKDIR /src

RUN apt-get update \
    && apt-get install -y --no-install-recommends curl make \
    && rm -rf /var/lib/apt/lists/*

RUN go install github.com/a-h/templ/cmd/templ@v0.2.747

ARG TARGETARCH
RUN case "$TARGETARCH" in \
        amd64) tailwind_arch=x64 ;; \
        arm64) tailwind_arch=arm64 ;; \
        *) echo "Unsupported architecture: $TARGETARCH" >&2; exit 1 ;; \
    esac \
    && curl -fsSL "https://github.com/tailwindlabs/tailwindcss/releases/download/v4.1.11/tailwindcss-linux-${tailwind_arch}" \
        -o /usr/local/bin/tailwindcss \
    && chmod +x /usr/local/bin/tailwindcss

COPY go.mod ./
RUN go mod download

COPY . .
RUN go mod tidy
RUN GOFLAGS=-trimpath CGO_ENABLED=1 make build tailwindcss \
    && install -Dm755 a.out /out/portfolio

FROM docker.io/library/debian:bookworm-slim

WORKDIR /app
COPY --from=build /out/portfolio /app/portfolio
COPY --from=build /src/static /app/static

ENV PORT=:4000
EXPOSE 4000
ENTRYPOINT ["/app/portfolio"]
