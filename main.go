package main

import (
	"database/sql"
	"fmt"
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

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	dbQueries := database.New(db)

	c, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	state := State{
		db:  dbQueries,
		cfg: &c,
	}

	commands := Commands{
		handlers: make(map[string]func(*State, Command) error),
	}
	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: not enough arguments. Please provide a command.")
		os.Exit(1)
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	cmd := Command{
		name: cmdName,
		args: cmdArgs,
	}

	err = commands.run(&state, cmd)
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
