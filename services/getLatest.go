package services

import (
	"context"
	"log"

	"github.com/jufianto/comic-info-scraper/tasks"
)

func (c *Client) GetHomeLatests() error {
	ctx := context.Background()

	log.Println("navigating to homepage ", c.URLsite)

	if err := RunWithDefaultTimeout(ctx, tasks.Navigate(c.URLsite)); err != nil {
		return ReturnErrors("navigate", err)
	}

	return nil
}
