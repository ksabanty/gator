package main

import (
	"errors"
)

type Command struct {
	name string
	args []string
}

type Commands struct {
	cmdMap map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, handler func(*State, Command) error) {
	c.cmdMap[name] = handler
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.cmdMap[cmd.name]
	if !ok {
		return errors.New("unknown command")
	}
	return handler(s, cmd)
}
