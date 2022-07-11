package tasks

import "github.com/chromedp/chromedp"

func Navigate(urlStr string) chromedp.Action {
	return navigateTo(urlStr)
}
