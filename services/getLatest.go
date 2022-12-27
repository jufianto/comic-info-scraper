package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/jufianto/comic-info-scraper/tasks"
)

func (c *Client) GetHomeLatests(ctx context.Context) error {

	log.Println("navigating to homepage ", c.URLsite)

	if err := RunWithDefaultTimeout(ctx, tasks.Navigate(c.URLsite)); err != nil {
		// return ReturnErrors("navigate", err)
		log.Println("failed to navigate")
	}

	log.Println("success navigating to homepages")

	nextPage := true
	pages := 1
	var results []InfoComic

	for nextPage {
		var allNodes []*cdp.Node
		if err := RunWithDefaultTimeout(ctx, tasks.GetAllNodesHome(&allNodes)); err != nil {
			return ReturnErrors("GetAllNodes", err)
		}

		fmt.Println("Have total ", len(allNodes))

		for key, nodes := range allNodes {
			var title, chapter string
			if err := RunWithDefaultTimeout(ctx, tasks.GetTitle(nodes, &title)); err != nil {
				return ReturnErrors("getTitle", err)
			}
			if err := RunWithDefaultTimeout(ctx, tasks.GetChapter(nodes, &chapter)); err != nil {
				ReturnErrors("getChapter", err)
			}
			log.Printf("Got Title %d: %s Last Chapter: %s ", key+1, title, chapter)
			result := InfoComic{
				Title:       title,
				LastChapter: chapter,
			}
			results = append(results, result)
		}
		log.Printf("success getting %d pages \n", pages)

		log.Println("check if next pages exists")
		var nodesNext []*cdp.Node
		if err := RunWithDefaultTimeout(ctx, tasks.CheckNextPages(&nodesNext)); err != nil {
			return ReturnErrors("CheckNextPages", err)
		}
		if len(nodesNext) > 0 {
			nextPage = true
			if err := RunWithDefaultTimeout(ctx, tasks.ClickNextPages()); err != nil {
				return ReturnErrors("ClickNextPages", err)
			}
		} else {
			nextPage = false
		}
	}

	fmt.Println("Have total comic", len(results))

	time.Sleep(time.Hour)

	return nil
}
