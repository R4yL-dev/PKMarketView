package itemsparser

import (
	"errors"
	"strconv"
	"strings"

	"github.com/R4yL-dev/PKMarketView/internal/seller"
	"github.com/go-rod/rod"
)

func buildSeller(content *rod.Element) (*seller.Seller, error) {
	name, err := getSellerName(content)
	if err != nil {
		return nil, err
	}
	level, err := getSellerLevel(content)
	if err != nil {
		return nil, err
	}
	nationality, err := getSellerNationality(content)
	if err != nil {
		return nil, err
	}
	return seller.NewSeller(
		name,
		level,
		nationality,
	), nil
}

func getSellerName(content *rod.Element) (string, error) {
	spans, err := content.Elements("span")
	if err != nil {
		return "", err
	}
	for _, span := range spans {
		link, err := span.Element("a")
		if err != nil {
			continue
		}
		name, err := link.Text()
		if err != nil {
			return "", err
		}
		return name, nil
	}
	return "", errors.New("seller name not found")
}

func getSellerLevel(content *rod.Element) (int, error) {
	spans, err := content.Elements("span")
	if err != nil {
		return -1, err
	}
	for _, span := range spans {
		if hasClass(span, "badge") {
			tooltipAttr, err := span.Attribute("data-bs-original-title")
			if err != nil {
				return -1, err
			}
			text := strings.ReplaceAll(*tooltipAttr, "\u00a0", " ")
			parts := strings.Split(text, " ")
			level, err := strconv.Atoi(parts[0])
			if err != nil {
				return -1, err
			}
			return level, nil
		}
	}
	return -1, errors.New("seller level not found")
}

func getSellerNationality(content *rod.Element) (string, error) {
	spans, err := content.Elements("span")
	if err != nil {
		return "", err
	}
	for _, span := range spans {
		_, err := span.Element("span")
		if err != nil {
			continue
		}
		label, err := span.Attribute("aria-label")
		if err != nil {
			return "", err
		}
		if label == nil {
			continue
		}
		parts := strings.Split(*label, ": ")
		if len(parts) < 2 {
			return "", errors.New("nationality unknown format")
		}
		return parts[1], nil
	}
	return "", errors.New("nationality not found")
}
