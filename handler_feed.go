package main

import (
	"context"
	"fmt"

	"github.com/DisguisedTrolley/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: cli addfeed <name> <url>")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error following feed: %v", err)
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
