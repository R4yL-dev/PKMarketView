package itemsparser

import (
	"errors"
	"strconv"
	"strings"

	"github.com/R4yL-dev/PKMarketView/internal/item"
	"github.com/R4yL-dev/PKMarketView/internal/price"
	"github.com/R4yL-dev/PKMarketView/internal/seller"
	"github.com/go-rod/rod"
)

func buildItem(content *rod.Element, seller *seller.Seller) (*item.Item, error) {
	price, err := getItemPrice(content)
	if err != nil {
		return nil, err
	}
	quantity, err := getItemQuantity(content)
	if err != nil {
		return nil, err
	}
	comment, err := getItemComment(content)
	if err != nil {
		return nil, err
	}
	lang, err := getItemLang(content)
	if err != nil {
		return nil, err
	}
	return item.NewItem(
		price,
		quantity,
		comment,
		lang,
		seller,
	), nil
}

func getItemPrice(content *rod.Element) (price.Price, error) {
	priceSpan, err := content.Element("span.color-primary.fw-bold")
	if err != nil {
		return 0, err
	}
	priceStr, err := priceSpan.Text()
	if err != nil {
		return 0, err
	}
	priceStr = strings.ReplaceAll(priceStr, "€", "")
	priceStr = strings.ReplaceAll(priceStr, ".", "")
	priceStr = strings.TrimSpace(priceStr)
	res, err := price.NewPrice(priceStr)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func getItemQuantity(content *rod.Element) (int, error) {
	quantitySpan, err := content.Element("span.item-count")
	if err != nil {
		return 0, err
	}
	quantityStr, err := quantitySpan.Text()
	if err != nil {
		return 0, err
	}
	res, err := strconv.Atoi(quantityStr)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func getItemComment(content *rod.Element) (string, error) {
	commentSpan, err := content.Element(".product-comments")
	if err != nil {
		return "", nil
	}
	commentStr, err := commentSpan.Text()
	if err != nil {
		return "", err
	}
	return commentStr, nil
}

func getItemLang(content *rod.Element) (string, error) {
	langSpan, err := content.Element("div.product-attributes span")
	if err != nil {
		return "", err
	}
	langStr, err := langSpan.Attribute("aria-label")
	if err != nil {
		return "", err
	}
	if langStr == nil {
		return "", errors.New("no language found")
	}
	return *langStr, nil
}
