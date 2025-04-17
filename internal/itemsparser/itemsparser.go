package itemsparser

import (
	"github.com/R4yL-dev/PKMarketView/internal/item"
	"github.com/go-rod/rod"
)

func Parse(content *rod.Element) ([]item.Item, error) {
	sellerInfo, err := extractSellerProductInfoDiv(content)
	if err != nil {
		return nil, err
	}
	items, err := buildItemsList(sellerInfo)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func extractSellerProductInfoDiv(content *rod.Element) (rod.Elements, error) {
	res, err := content.Elements(".col-sellerProductInfo")
	if err != nil {
		return rod.Elements{}, err
	}
	//Delete first element (table header)
	res = res[1:]
	return res, nil
}

func buildItemsList(content rod.Elements) ([]item.Item, error) {
	var items []item.Item
	for _, div := range content {
		sellerInfos, err := div.Element(".seller-name")
		if err != nil {
			return nil, err
		}
		seller, err := buildSeller(sellerInfos)
		if err != nil {
			return nil, err
		}
		itemInfo, err := div.Element(".col-product .row")
		if err != nil {
			return nil, err
		}
		item, err := buildItem(itemInfo, seller)
		if err != nil {
			return nil, err
		}
		items = append(items, *item)
	}
	return items, nil
}
