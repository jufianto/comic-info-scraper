package tasks

import (
	"context"
	"fmt"

	"github.com/chromedp/cdproto/browser"
	"github.com/chromedp/chromedp"
)

func navigateTo(url string) chromedp.Action {
	return chromedp.Navigate(url)
}

func ResizeWindow(width, height int64) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		var id browser.WindowID
		var bounds *browser.Bounds

		var err error
		action := chromedp.ActionFunc(func(ctx context.Context) error {
			id, bounds, err = browser.GetWindowForTarget().Do(ctx)
			return err
		})

		if err := chromedp.Run(ctx, action); err != nil {
			return fmt.Errorf("could not get target: %w", err)
		}

		bounds.Width = width
		bounds.Height = height

		if err := chromedp.Run(ctx, browser.SetWindowBounds(id, bounds)); err != nil {
			return fmt.Errorf("could not set window bounds: %w", err)
		}
		return nil
	})
}
