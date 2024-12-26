package tasks

import (
	"fmt"

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
	return chromedp.Nodes(`//h3[contains(text(), "Latest Update")]//parent::div/following-sibling::div[1]/div`, nodes)
}

func GetTitle(index int, result *string) chromedp.Action {
	sel := fmt.Sprintf(`(//h3[contains(text(), "Latest Update")]//parent::div/following-sibling::div[1]/div//span/a)[%d]`, index)
	fmt.Println("sel", sel)
	// Use the correct query selector to target the title based on the structure
	return chromedp.Text(sel, result, chromedp.NodeReady, chromedp.NodeVisible)
}

func GetChapter(index int, result *string) chromedp.Action {
	sel := fmt.Sprintf(`(//h3[contains(text(), "Latest Update")]/parent::div/following-sibling::div[1]/div[%d]//p[contains(text(), "Chapter")])[1]`, index)
	fmt.Println("sel", sel)
	return chromedp.Text(sel, result, chromedp.NodeReady, chromedp.NodeVisible)
}

func CheckNextPages(nodeNext *[]*cdp.Node) chromedp.Action {
	return chromedp.Nodes(`//a[contains(text(), "Next")]`, nodeNext, chromedp.AtLeast(0))
}

func ClickNextPages() chromedp.Action {
	return chromedp.Click(`//a[contains(text(), "Next")]`)
}

func GetAttribute(node *cdp.Node, attr string, result *string, ok *bool) chromedp.Action {
	return chromedp.AttributeValue(`//a[contains(text(), "Next")]`, attr, result, ok)
}
