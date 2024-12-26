package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/jufianto/comic-info-scraper/tasks"
)

// TODO: store cache and cookies for faster access
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
		fmt.Println("get all nodes home")
		if err := RunWithDefaultTimeout(ctx, tasks.GetAllNodesHome(&allNodes)); err != nil {
			return ReturnErrors("GetAllNodes", err)
		}

		fmt.Println("Have total ", len(allNodes))

		for key, _ := range allNodes {
			var title, chapter string

			fmt.Println("key", key+1)

			if err := RunWithDefaultTimeout(ctx, tasks.GetTitle(key+1, &title)); err != nil {
				return ReturnErrors("getTitle", err)
			}
			if err := RunWithDefaultTimeout(ctx, tasks.GetChapter(key+1, &chapter)); err != nil {
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
		if err := RunWithTimeout(ctx, 5*time.Second, tasks.CheckNextPages(&nodesNext)); err != nil {
			return ReturnErrors("CheckNextPages", err)
		}
		if len(nodesNext) > 0 {

			log.Println("found node next", len(nodesNext))
			// get attribute first
			var attr string
			var ok bool
			if err := RunWithDefaultTimeout(ctx, tasks.GetAttribute(nodesNext[0], "href", &attr, &ok)); err != nil {
				log.Println("failed tp get attribute", err)
				if len(results) > 0 {
					// if already have the result, return immediately
					return nil
				}
				return ReturnErrors("GetAttribute", err)
			}

			fmt.Println("get attribute href", attr)

			if attr == "#" {
				nextPage = false
				break
			}

			nextPage = true
			time.Sleep(2 * time.Second)
			if err := RunWithDefaultTimeout(ctx, tasks.ClickNextPages()); err != nil {
				return ReturnErrors("ClickNextPages", err)
			}
		} else {
			nextPage = false
		}
	}

	fmt.Println("Have total comic", len(results))
	for _, res := range results {
		fmt.Println("Title", res.Title, "Last Chapter", res.LastChapter)
	}

	time.Sleep(time.Hour)

	return nil
}
