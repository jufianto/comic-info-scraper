package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

var defaultTimeout = 2 * time.Minute
var errorMsgFr = "failed to execute tasks %s: %v"

func RunWithDefaultTimeout(ctx context.Context, actions ...chromedp.Action) error {
	ctxs, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	return chromedp.Run(ctxs, actions...)
}

func ReturnErrors(tasks string, err error) error {
	errMsg := fmt.Sprintf(errorMsgFr, tasks, err)
	log.Println(errMsg)
	return fmt.Errorf(errMsg)
}
