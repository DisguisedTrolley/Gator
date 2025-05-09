package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/DisguisedTrolley/gator/internal/database"
	"github.com/lib/pq"
)

const LAYOUT = "Mon, 02 Jan 2006 15:04:05 -0700"

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
		pubTime, err := time.Parse(LAYOUT, item.PubDate)
		if err != nil {
			log.Fatal("time format incorrect: ", err)
		}

		desc := sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			Title:       item.Title,
			Description: desc,
			Url:         item.Link,
			PublishedAt: pubTime,
			FeedID:      feed.ID,
		})

		if err == nil {
			continue
		}

		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code != "23505" {
				log.Fatal("error creating post: ", pgErr.Error())
			}
		}

		fmt.Println("Added post: ", item.Title)

	}

	fmt.Println()
	fmt.Println("===================================")
	fmt.Println()
}
