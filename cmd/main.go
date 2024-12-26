package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"

	"github.com/jufianto/comic-info-scraper/cmd/config"
	cl "github.com/jufianto/comic-info-scraper/services"
)

func main() {
	// set config
	log.Println("getting config from config file yaml")
	cfgs, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config %s", err)
		return
	}

	urlSite := cfgs.GetString("url")

	client := cl.NewClient(urlSite, cl.WithHeadfull(true))

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()
	userDataDir := "./userdir"

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.WindowSize(1920, 1080),
		// chromedp.ExecPath("/home/sv/chrome-linux/chrome"),
		chromedp.UserDataDir(userDataDir),
	)

	if !cfgs.GetBool("headless") {
		opts = append(opts, chromedp.Flag("headless", false))
	}

	alloCtx, cancelAllCtx := chromedp.NewExecAllocator(ctx, opts...)
	defer cancelAllCtx()

	cdpCtx, cancel := chromedp.NewContext(alloCtx)
	defer cancel()

	log.Printf("will getting data from site %s", client.URLsite)

	if err := chromedp.Run(cdpCtx); err != nil {
		log.Fatalf("failed to start first browser %v", err)
	}

	time.AfterFunc(time.Hour, func() {
		fmt.Println("force cancel after 1 minutes browser")
		cancel()
	})

	log.Println("emulate viewport")
	if err := chromedp.Run(cdpCtx,
		chromedp.Emulate(device.Reset),
		chromedp.EmulateViewport(1336, 768),
	); err != nil {
		log.Printf("failed to emulate viewport: %v \n", err)
	}

	if err := client.GetHomeLatests(cdpCtx); err != nil {
		log.Println("failed to run get home services", err)
	}
}
