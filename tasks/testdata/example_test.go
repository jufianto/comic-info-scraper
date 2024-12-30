package testdata

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func Test_Example(t *testing.T) {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.WindowSize(1920, 1080),
		chromedp.Flag("headless", false),
	)

	alloCtx, cancelAllCtx := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAllCtx()

	cdpCtx, cancelChrome := chromedp.NewContext(alloCtx)
	defer cancelChrome()

	if err := chromedp.Run(cdpCtx,
		chromedp.Navigate("http://127.0.0.1:5500/tasks/testdata/index.html")); err != nil {
		t.Fatal(err)
	}

	var mainNode []*cdp.Node
	if err := chromedp.Run(cdpCtx,
		chromedp.Nodes(`//div[@class="ex1"]`, &mainNode),
	); err != nil {
		t.Log(err)
	}

	var text string
	if err := chromedp.Run(cdpCtx,
		chromedp.Text(`span>a`, &text, chromedp.ByQueryAll, chromedp.FromNode(mainNode[0])),
	); err != nil {
		fmt.Println("Error", err)
		t.Log(err)
	}

	fmt.Println("text", text)

	time.Sleep(time.Hour)

}
