package main

import (
	"context"
	"fmt"

	"github.com/DisguisedTrolley/gator/internal/database"
)

func middlewareLoggedIn(
	handler func(s *state, cmd command, user database.User) error,
) func(*state, command) error {
	return func(s *state, c command) error {
		user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("user not found")
		}

		return handler(s, c, user)
	}
}
