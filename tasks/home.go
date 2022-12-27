package tasks

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func Navigate(urlStr string) chromedp.Action {
	return navigateTo(urlStr)
}

func GetTestAction() chromedp.Action {
	return chromedp.WaitVisible("#jufi", chromedp.ByID)
}

func GetAllNodesHome(nodes *[]*cdp.Node) chromedp.Action {
	return chromedp.Nodes(`//div[@class="utao styletwo"]`, nodes)
}

func GetTitle(nodes *cdp.Node, result *string) chromedp.Action {
	return chromedp.Text(`h4`, result, chromedp.ByQuery, chromedp.FromNode(nodes))
}

func GetChapter(nodes *cdp.Node, result *string) chromedp.Action {
	return chromedp.Text(`ul > li:nth-child(1) > a`, result, chromedp.ByQueryAll, chromedp.FromNode(nodes))
}

func CheckNextPages(nodeNext *[]*cdp.Node) chromedp.Action {
	return chromedp.Nodes(`//a[contains(text(), "Next")]`, nodeNext, chromedp.AtLeast(0))
}

func ClickNextPages() chromedp.Action {
	return chromedp.Click(`//a[contains(text(), "Next")]`)
}
