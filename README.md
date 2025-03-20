# Gator

A multi-player command line tool for aggregating RSS feeds and viewing the posts.

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `gator` with:

```bash
go install ...
```

## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## Usage

Create a new user:

```bash
blog-aggregator register <name>
```

Add a feed:

```bash
blog-aggregator addfeed <url>
```

Start the aggregator:

```bash
blog-aggregator agg 30s
```

View the posts:

```bash
blog-aggregator browse [limit]
```

There are a few other commands you'll need as well:

- `blog-aggregator login <name>` - Log in as a user that already exists
- `blog-aggregator users` - List all users
- `blog-aggregator feeds` - List all feeds
- `blog-aggregator follow <url>` - Follow a feed that already exists in the database
- `blog-aggregator unfollow <url>` - Unfollow a feed that already exists in the database
