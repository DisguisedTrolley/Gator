package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}

	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("New user set successfully.")

	return nil
}
