package item

import (
	"fmt"

	"github.com/R4yL-dev/PKMarketView/internal/price"
	"github.com/R4yL-dev/PKMarketView/internal/seller"
)

func (item *Item) GetPrice() price.Price {
	return item.price
}

func (item *Item) GetQuantity() int {
	return item.quantity
}

func (item *Item) GetComment() string {
	return item.comment
}

func (item *Item) GetLang() string {
	return item.lang
}

func (item *Item) GetSeller() *seller.Seller {
	return item.seller
}

func (item *Item) SetPrice(price price.Price) {
	item.price = price
}

func (item *Item) SetQuantity(quantity int) {
	item.quantity = quantity
}

func (item *Item) SetComment(comment string) {
	item.comment = comment
}

func (item *Item) SetLang(lang string) {
	item.lang = lang
}

func (item *Item) SetSeller(seller *seller.Seller) {
	item.seller = seller
}

func (item *Item) Show() {
	fmt.Printf("item price: %s\n", item.price)
	fmt.Printf("item quantity: %d\n", item.quantity)
	fmt.Printf("item comment: %s\n", item.comment)
	fmt.Printf("item lang: %s\n", item.lang)
	item.seller.Show()
}
