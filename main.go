package main

import (
	"fmt"
	"os"

	"github.com/exglegaming/blog-aggregator/internal/config"
)

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

type State struct {
	cfg *config.Config
}

type Command struct {
	name string
	args []string
}

type Commands struct {
	handlers map[string]func(*State, Command) error
}

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("A user name must be entered")
	}

	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Username has been set to: %s\n", username)
	return nil
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	if c.handlers == nil {
		c.handlers = make(map[string]func(*State, Command) error)
	}
	c.handlers[name] = f
}

func (c *Commands) run(s *State, cmd Command) error {
	command, exists := c.handlers[cmd.name]
	if !exists {
		return fmt.Errorf("Command %s not found", cmd.name)
	}

	err := command(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
