package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}

	name := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	err = s.config.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Println("New user set successfully.")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the register handler expects a single argument, the username")
	}

	name := cmd.args[0]

	_, err := s.db.CreateUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("user already exists")
	}

	err = s.config.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Println("New user created successfully.")

	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to fetch users: %v", err)
	}

	for _, user := range users {
		str := fmt.Sprintf("* %s", user.Name)
		if user.Name == s.config.CurrentUserName {
			str += " (current)"
		}

		fmt.Println(str)
	}

	return nil
}

func handlerReset(s *state, _ command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting database")
	}

	fmt.Println("Database reset successful")

	return nil
}
