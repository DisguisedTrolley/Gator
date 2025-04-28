package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: cli agg <time_between_reqs>")
	}

	timeBwReqs := cmd.args[0]

	duration, err := time.ParseDuration(timeBwReqs)
	if err != nil {
		return fmt.Errorf("improper duration")
	}

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("No new feeds")
		return
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Fatal("unable to mark feed as fetched")
	}

	fetched, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range fetched.Channel.Item {
		fmt.Println(item.Title)
	}

	fmt.Println()
	fmt.Println("=================================")
	fmt.Println("=================================")
	fmt.Println()
}
