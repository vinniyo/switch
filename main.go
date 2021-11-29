package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	x := 1
	for {
		check("https://store.nintendo.com/nintendo-switch-oled-model-white-set.html", x)
		check("https://store.nintendo.com/nintendo-switch-oled-model-neon-blue-neon-red-set.html", x)
		x++

		if x%100 == 0 {
			fmt.Println("deleting Temp files")
			go deleteAll()
		}
	}
}

func deleteAll() {
	//ChromeDP creates temp files that need to be cleared out. Be careful
	userDir, _ := os.UserCacheDir()
	os.RemoveAll(userDir + "/Temp")

}

func check(page string, x int) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36"),
		chromedp.Flag("enable-automation", false),
		//To view what it's doing // chromedp.Flag("headless", false),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(
		ctx,
	)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var datar string

	err := chromedp.Run(ctx,

		chromedp.Navigate(page),
		chromedp.InnerHTML(`#product-addtocart-button > span`, &datar, chromedp.BySearch),
	)
	if err != nil {
		//this just errors out if the Add to Cart button isn't found.
		fmt.Println("Tried: ", page, " iteration: ", x)
		return

	}

	fmt.Println("dater: ", datar)
	if datar == "Add To Cart" {

		fmt.Println("IN STOCK")

		addCart(page)

		time.Sleep(400 * time.Minute)
	} else {
		fmt.Println("Not in stock")
	}

}

func addCart(page string) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36"),
		chromedp.Flag("headless", false),
		chromedp.Flag("enable-automation", false),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(
		ctx,
	)
	defer cancel()

	err := chromedp.Run(ctx,

		chromedp.Navigate(page),
		chromedp.Click("#product-addtocart-button > span", chromedp.ByID),
		chromedp.WaitVisible("button.myButtonthatwillneverexist"),
	)
	if err != nil {
		fmt.Println("Here is the error: ", err)
	}

}
