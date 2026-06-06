FROM golang:1.22-alpine AS build

# sqlite3 requires CGO and gcc
RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o server .

# ── Runtime ──────────────────────────────────────────────────────────────────
FROM alpine:3.19

RUN apk add --no-cache sqlite

WORKDIR /app
COPY --from=build /app/server .

RUN mkdir -p /data

ENV DB_PATH=/data/shop.db
ENV PORT=8080

EXPOSE 8080
CMD ["./server"]
