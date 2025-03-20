package main

import (
	"database/sql"
	"github.com/exglegaming/blog-aggregator/internal/config"
	"github.com/exglegaming/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type State struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	state := &State{
		db:  dbQueries,
		cfg: &cfg,
	}

	commands := Commands{
		handlers: make(map[string]func(*State, Command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerGetUsers)
	commands.register("agg", handlerAgg)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = commands.run(state, Command{name: cmdName, args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
