package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/R4yL-dev/PKMarketView/internal/scrapper"
)

func main() {
	urlStr := "https://www.cardmarket.com/fr/Pokemon/Products/Elite-Trainer-Boxes/Prismatic-Evolutions-10-Elite-Trainer-Box-Case"
	url, err := url.Parse(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	items, err := scrapper.GetItemList(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		item.Show()
		fmt.Println("--------------------")
	}
}
