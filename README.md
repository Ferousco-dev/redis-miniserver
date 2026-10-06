# Redis Mini Server

A minimal Redis cache wrapper written in Go using [`go-redis/v9`](https://github.com/redis/go-redis).

It provides a simple `Cache` abstraction for connecting to Redis, checking connectivity, storing values, and retrieving them with optional TTL support.

## Features

- Redis connection and health check
- Key-value storage and retrieval
- TTL support
- Lightweight cache abstraction
- Built with Go and `go-redis/v9`

## Requirements

- Go
- Redis running on `localhost:6379`

## Getting Started

Clone the repository:

```bash
git clone https://github.com/Ferousco-dev/redis-miniserver.git
cd redis-miniserver
```

Install dependencies:

```bash
go mod download
```

Make sure Redis is running:

```bash
redis-cli ping
```

Expected output:

```
PONG
```

Run the project:

```bash
go run .
```

## Project Structure

```
redis-miniserver/
├── cache/
│   └── cache.go
├── main.go
├── go.mod
└── go.sum
```

## Usage

```go
cache := cache.New("localhost:6379", "", 0)

cache.Set(ctx, "user:name", "feranmi", 0)

value, _ := cache.Get(ctx, "user:name")

fmt.Println(value)
```

The final argument of `Set` is the TTL:

```go
cache.Set(ctx, "session:user", "active", 5*time.Minute)
```

Use `0` for a value without an expiration.

## Development

Run tests:

```bash
go test ./...
```

Run static checks:

```bash
go vet ./...
```

Build:

```bash
go build .
```

Format:

```bash
gofmt -w .
```

## Tech Stack

- Go
- Redis
- go-redis/v9

> A small project for learning Redis, Go caching patterns, and backend infrastructure.
