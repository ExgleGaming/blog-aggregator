package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *State, cmd Command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", feed)

	return nil
}
