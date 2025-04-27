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

func handlerFollowFeed(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: cli addfeed <name> <url>")
	}

	url := cmd.args[0]
	feedID, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("feed with given url doesn't exist")
	}

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feedID,
	})
	if err != nil {
		return fmt.Errorf("error following feed: %v", err)
	}

	fmt.Println("followed feed: ", url)

	return nil
}

func handlerListFollowingFeeds(s *state, _ command) error {
	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("not following any feed")
	}

	for _, feed := range feeds {
		fmt.Println("* ", feed.FeedName)
	}

	return nil
}
