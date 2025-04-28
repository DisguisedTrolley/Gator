package main

import (
	"context"
	"fmt"

	"github.com/DisguisedTrolley/gator/internal/database"
)

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: cli addfeed <name> <url>")
	}

	url := cmd.args[0]
	feedID, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("feed with given url doesn't exist")
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

func handlerListFollowingFeeds(s *state, _ command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("not following any feed")
	}

	for _, feed := range feeds {
		fmt.Println("* ", feed.FeedName)
	}

	return nil
}

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: cli unfollow <url>")
	}

	url := cmd.args[0]

	err := s.db.UnfollowFeedForUser(context.Background(), database.UnfollowFeedForUserParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %v", err)
	}

	fmt.Println("unfollowed feed: ", url)

	return nil
}
