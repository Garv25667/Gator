# Gator

A multi-user RSS feed aggregator CLI built in Go, backed by PostgreSQL.

## Prerequisites

- [Go 1.22+](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation

```bash
go install github.com/yourusername/gator@latest
```

## Config Setup

Create a config file at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

## Database Setup

```bash
goose postgres "postgres://username:password@localhost:5432/gator" up
```

## Commands

```bash
# Register a new user
gator register <name>

# Login as existing user
gator login <name>

# Add an RSS feed
gator addfeed <name> <url>

# Follow an existing feed
gator follow <url>

# Start aggregating (fetches every 30s)
gator agg 30s

# Browse latest posts
gator browse 10
```