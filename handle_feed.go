package main

import (
	"context"
	"fmt"
	"github.com/exglegaming/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"time"
)

func printFeed(feed *database.Feed, user *database.User) {
	fmt.Printf("Feed Name: %s\n", feed.Name)
	fmt.Printf("Feed URL: %s\n", feed.Url)
	fmt.Printf("Feed User: %s\n", user.Name)
}

func handlerAddFeed(s *State, cmd Command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	name, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	feedName := cmd.args[0]
	feedUrl := cmd.args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    name.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Printf("Feed: %v\n", feed)
	return nil
}

func handlerGetFeed(s *State, cmd Command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user with id: %v, error: %w", feed.UserID, err)
		}
		printFeed(&feed, &user)
	}
	return nil
}
