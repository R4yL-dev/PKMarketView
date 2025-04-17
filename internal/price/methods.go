package price

import "fmt"

func (price Price) String() string {
	euros := int(price) / 100
	centimes := int(price) % 100
	return fmt.Sprintf("%d.%02d €", euros, centimes)
}
