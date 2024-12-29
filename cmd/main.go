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
	"github.com/jufianto/comic-info-scraper/store"
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

	log.Println("init store")
	fs, closeFunc, err := store.InitStore(ctx, cfgs.GetString("firestore.project_id"), cfgs.GetString("firestore.service_account_credential"))
	defer closeFunc()
	if err != nil {
		log.Fatalf("failed to init store %v", err)
		return
	}

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

	cdpCtx, cancelChrome := chromedp.NewContext(alloCtx)
	defer cancelChrome()

	log.Printf("will getting data from site %s", client.URLsite)

	if err := chromedp.Run(cdpCtx); err != nil {
		log.Fatalf("failed to start first browser %v", err)
	}

	time.AfterFunc(time.Hour, func() {
		fmt.Println("force cancel after 1 minutes browser")
		cancelChrome()
	})

	log.Println("emulate viewport")
	if err := chromedp.Run(cdpCtx,
		chromedp.Emulate(device.Reset),
		chromedp.EmulateViewport(1336, 768),
	); err != nil {
		log.Printf("failed to emulate viewport: %v \n", err)
	}

	results, err := client.GetHomeLatests(cdpCtx)
	if err != nil {
		log.Println("failed to run get home services", err)
		return
	}

	log.Printf("success get the results, total results %d \n", len(results))

	_, err = store.StoreToYaml(results)
	if err != nil {
		log.Printf("failed store to yaml file: %v \n", err)
	}

	log.Println("store to yaml success, try store to firestore")
	mapStrData, err := store.ConvertToJSON(results)
	if err != nil {
		log.Fatalf("failed to convert to json", err)
		return
	}
	fmt.Println("xx", mapStrData)
	err = fs.StoreComic(ctx, cfgs.GetString("firestore.collection"), mapStrData)
	if err != nil {
		log.Fatalf("failed to store data to firestore %v", err)
		return
	}
	log.Println("sleep for 1 minutes")
	time.Sleep(time.Minute)

	log.Println("shutdown browser")
	cancelChrome()
}
