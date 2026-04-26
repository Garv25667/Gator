package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/bootdotdev/gator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("error getting feed to fetch: %v", err)
		return
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Printf("Error in marking feed fetched: %v", err)
		return
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url.String)
	if err != nil {
		fmt.Printf("error fetching feed: %v", err)
		return
	}
	for _, item := range rssFeed.Channel.Item {
		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			publishedAt, err = time.Parse(time.RFC1123, item.PubDate)
			if err != nil {
				publishedAt = time.Time{}
			}
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         sql.NullString{String: item.Link, Valid: item.Link != ""},
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			// ignore duplicate URL errors, log everything else
			if !strings.Contains(err.Error(), "duplicate key") {
				fmt.Printf("couldn't save post: %v\n", err)
			}
		}

	}

}
