package main

import (
	"context"
	"fmt"
	"github.com/exglegaming/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"time"
)

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
	return nil
}
