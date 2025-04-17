package price

import (
	"fmt"
	"strconv"
	"strings"
)

type Price int

func NewPrice(str string) (Price, error) {
	str = strings.ReplaceAll(str, ",", ".")

	parts := strings.SplitN(str, ".", 2)
	euros := 0
	centimes := 0
	var err error

	if len(parts) > 0 {
		euros, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, fmt.Errorf("euros unvalid: %w", err)
		}
	}

	if len(parts) == 2 {
		centStr := parts[1]
		if len(centStr) > 2 {
			centStr = centStr[:2]
		} else if len(centStr) < 2 {
			centStr += strings.Repeat("0", 2-len(centStr))
		}

		centimes, err = strconv.Atoi(centStr)
		if err != nil {
			return 0, fmt.Errorf("centimes unvalid: %w", err)
		}
	}
	return Price(euros*100 + centimes), nil
}
