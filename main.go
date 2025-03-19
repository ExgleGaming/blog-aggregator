package main

import (
	"fmt"
	"os"

	"github.com/exglegaming/blog-aggregator/internal/config"
)

type State struct {
	cfg *config.Config
}

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	state := State{
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
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

}
