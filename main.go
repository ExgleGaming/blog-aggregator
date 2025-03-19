package main

import (
	"fmt"
	"os"

	"github.com/exglegaming/blog-aggregator/internal/config"
)

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	err = c.SetUser("Hunter")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("DB URL: ", cfg.DbURL)
	fmt.Println("Current User: ", cfg.CurrentUserName)
}
