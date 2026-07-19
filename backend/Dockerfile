FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o vidrop ./cmd/api


FROM python:3.13-slim

WORKDIR /app

RUN apt-get update && apt-get install -y \
    ffmpeg \
    curl \
    unzip \
    && rm -rf /var/lib/apt/lists/*

RUN curl -fsSL https://deno.land/install.sh | sh

ENV PATH="/root/.deno/bin:$PATH"

RUN pip install --no-cache-dir -U \
    yt-dlp \
    yt-dlp-ejs

COPY --from=builder /app/vidrop .

EXPOSE 8080

CMD ["./vidrop"]
