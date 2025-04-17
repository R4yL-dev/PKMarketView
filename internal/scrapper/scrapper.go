package scrapper

import (
	"errors"
	"net/url"
	"time"

	"github.com/R4yL-dev/PKMarketView/internal/item"
	"github.com/R4yL-dev/PKMarketView/internal/itemsparser"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

func init() {
	launcher.NewBrowser().MustGet()
}

// TODO: Mettre la logique de chargement de la page dans sa propre fonction
// TODO: Mettre la logique d extraction du contenu dans sa propre fonction
func GetItemList(url *url.URL) ([]item.Item, error) {
	browser := rod.New().Timeout(time.Minute).MustConnect()
	defer browser.MustClose()

	page, err := stealth.Page(browser)
	if err != nil {
		return nil, err
	}
	err = page.Navigate(url.String())
	if err != nil {
		return nil, err
	}
	err = page.WaitLoad()
	if err != nil {
		return nil, err
	}

	if !isValidProductPage(page) {
		return nil, errors.New("page is not a product page")
	}
	table, err := page.Element(".article-table")
	if err != nil {
		return nil, err
	}
	items, err := itemsparser.Parse(table)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func isValidProductPage(page *rod.Page) bool {
	hasArticleTable := page.MustHas(".article-table")
	hasNoResults := page.MustHas(".noResults")
	if !hasArticleTable {
		return false
	}
	if hasNoResults {
		return false
	}
	return true
}
