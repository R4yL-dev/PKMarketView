package itemsparser

import (
	"slices"
	"strings"

	"github.com/go-rod/rod"
)

func hasClass(el *rod.Element, className string) bool {
	classAttr := el.MustAttribute("class")
	if classAttr == nil {
		return false
	}
	classes := strings.Fields(*classAttr)
	return slices.Contains(classes, className)
}
