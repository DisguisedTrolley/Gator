package main

import (
	"context"
	"fmt"

	"github.com/DisguisedTrolley/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: cli addfeed <name> <url>")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	_, err = s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	fmt.Println("Feed created successfully")
	return nil
}

func handlerListFeed(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feeds")
	}

	for _, f := range feeds {
		fmt.Println("Name: ", f.FeedName)
		fmt.Println("URL: ", f.Url)
		fmt.Println("User: ", f.UserName)
		fmt.Println("==============================")
		fmt.Println()
	}

	return nil
}
