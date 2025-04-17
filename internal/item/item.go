package item

import (
	"github.com/R4yL-dev/PKMarketView/internal/price"
	"github.com/R4yL-dev/PKMarketView/internal/seller"
)

type Item struct {
	price    price.Price
	quantity int
	comment  string
	lang     string
	seller   *seller.Seller
}

func NewItem(price price.Price, quantity int, comment string, lang string, seller *seller.Seller) *Item {
	return &Item{
		price:    price,
		quantity: quantity,
		comment:  comment,
		lang:     lang,
		seller:   seller,
	}
}
