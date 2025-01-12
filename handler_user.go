package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ksabanty/gator/internal/database"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}
	name := cmd.args[0]

	user, notFound := s.db.GetUser(context.Background(), name)
	if notFound != nil {
		return fmt.Errorf("user not found")
	}

	err := s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}
	name := cmd.args[0]

	ctx := context.Background()
	params := database.CreateUserParams{ID: uuid.New(), Name: name}
	_, err := s.db.CreateUser(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	fmt.Printf("\nUser created successfully: %s", name)

	setErr := s.cfg.SetUser(name)
	if setErr != nil {
		return fmt.Errorf("couldn't set current user: %w", setErr)
	}

	fmt.Println("User switched successfully!")

	return nil
}

func handlerReset(s *State, cmd Command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete all users: %w", err)
	}

	fmt.Println("All users deleted successfully!")
	return nil
}

func handlerUsers(s *State, cmd Command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get all users: %w", err)
	}

	for _, user := range users {
		if user == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}
	}

	return nil
}
