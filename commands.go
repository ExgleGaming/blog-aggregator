package main

import (
	"errors"
)

type Command struct {
	name string
	args []string
}

type Commands struct {
	handlers map[string]func(*State, Command) error
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	if c.handlers == nil {
		c.handlers = make(map[string]func(*State, Command) error)
	}
	c.handlers[name] = f
}

func (c *Commands) run(s *State, cmd Command) error {
	f, ok := c.handlers[cmd.name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
