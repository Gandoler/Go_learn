FROM golang:1.22 AS build

WORKDIR /src


COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN go build -o /app/url_shortener ./cmd/url_shortener


FROM debian:bookworm-slim AS final

WORKDIR /app

COPY --from=build /app/url_shortener .


COPY ./config ./config

ENV CONFIG_PATH=/app/config/config.yaml


EXPOSE 8082

ENTRYPOINT ["./url_shortener"]
