FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o vidrop ./cmd/api


FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache \
    ffmpeg \
    yt-dlp \
    deno

COPY --from=builder /app/vidrop .

EXPOSE 8080

CMD ["./vidrop"]
