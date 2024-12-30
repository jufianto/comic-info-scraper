package services

import (
	"context"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/jufianto/comic-info-scraper/tasks"
)

// TODO: store cache and cookies for faster access
func (c *Client) GetHomeLatests(ctx context.Context) (results []InfoComic, err error) {

	log.Println("navigating to homepage ", c.URLsite)

	if err := RunWithDefaultTimeout(ctx, tasks.Navigate(c.URLsite)); err != nil {
		// return SetError("navigate", err)
		log.Println("failed to navigate")
	}

	log.Println("success navigating to homepages")

	nextPage := true
	pages := 1

	for nextPage {
		var allNodes []*cdp.Node
		log.Println("get all nodes home")

		if pages > 1 {
			time.Sleep(3 * time.Second) // to wait page full load. TODO: find another by using chromedp listen target
		}

		if err := RunWithDefaultTimeout(ctx, tasks.GetAllNodesHome(&allNodes)); err != nil {
			return nil, SetError("GetAllNodes", err)
		}

		log.Println("Have total ", len(allNodes))
		reNumber := regexp.MustCompile(`\d+`)
		reRemoveNumber := regexp.MustCompile(`\b\w*\d\w*\b`)
		rpl := strings.NewReplacer("-", " ", "/series/", "")

		for key := range allNodes {
			var title, chapter string

			if err := RunWithDefaultTimeout(ctx, tasks.GetTitleAttributeHref(key+1, &title)); err != nil {
				return nil, SetError("getTitle", err)
			}
			if err := RunWithDefaultTimeout(ctx, tasks.GetChapter(key+1, &chapter)); err != nil {
					SetError("getChapter", err)
			}

			title = rpl.Replace(title)

			title = reRemoveNumber.ReplaceAllString(title, "")
			title = strings.TrimSpace(title)

			chapter = reNumber.FindString(chapter)

			// log.Printf("Got Title %d: %s Last Chapter: %s ", key+1, title, chapter)
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
			return nil, SetError("CheckNextPages", err)
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
					return results, nil
				}
				return nil, SetError("GetAttribute", err)
			}

			if attr == "#" {
				nextPage = false
				break
			}

			nextPage = true
			time.Sleep(2 * time.Second)
			if err := RunWithDefaultTimeout(ctx, tasks.ClickNextPages()); err != nil {
				return nil, SetError("ClickNextPages", err)
			}
			pages++
		} else {
			nextPage = false
		}
	}

	log.Println("Have total comic", len(results))
	return results, nil
}
