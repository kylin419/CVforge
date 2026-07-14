package morphology

import (
	"fmt"
	"strings"
)

func ParseElement(name string) (Element, error) {
	switch strings.ToLower(name) {
	case "square":
		return Square3x3, nil
	case "cross":
		return Cross3x3, nil
	default:
		return Element{}, fmt.Errorf("unknown element: %s", name)
	}
}
