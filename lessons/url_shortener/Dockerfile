FROM golang:1.25 AS build

WORKDIR /src


RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

RUN git clone https://github.com/Gandoler/Go_learn.git

WORKDIR /src/Go_learn/lessons/url_shortener



RUN go mod download
RUN go build -o /app/url_shortener ./cmd/url_shortener


FROM debian:bookworm-slim AS final

WORKDIR /app

COPY --from=build /app/url_shortener .



ENV CONFIG_PATH=./cmd/config/prod.yaml


EXPOSE 8082

ENTRYPOINT ["./url_shortener"]
