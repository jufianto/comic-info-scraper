package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

var defaultTimeout = time.Minute
var errorMsgFr = "failed to execute tasks %s: %w"

func RunWithDefaultTimeout(ctx context.Context, actions ...chromedp.Action) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	return chromedp.Run(ctx, actions...)
}

func ReturnErrors(tasks string, err error) error {
	errMsg := fmt.Sprintf(errorMsgFr, tasks, err)
	log.Println(errMsg)
	return fmt.Errorf(errMsg)
}
