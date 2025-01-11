package main

import (
	"errors"
)

type Command struct {
	name string
	args []string
}

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return errors.New("must provide a username")
	}
	return s.cfg.SetUser(cmd.args[0])
}
