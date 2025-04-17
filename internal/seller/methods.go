package seller

import "fmt"

func (seller *Seller) GetName() string {
	return seller.name
}

func (seller *Seller) GetLevel() int {
	return seller.level
}

func (seller *Seller) GetNationality() string {
	return seller.nationality
}

func (seller *Seller) SetName(name string) {
	seller.name = name
}

func (seller *Seller) SetLevel(level int) {
	seller.level = level
}

func (seller *Seller) SetNationality(nationality string) {
	seller.nationality = nationality
}

func (seller *Seller) Show() {
	fmt.Printf("seller name: %s\n", seller.name)
	fmt.Printf("seller level: %d\n", seller.level)
	fmt.Printf("seller nationality: %s\n", seller.nationality)
}
