package tasks

import (
	"github.com/chromedp/chromedp"
)

func navigateTo(url string) chromedp.Action {
	return chromedp.Navigate(url)
}
