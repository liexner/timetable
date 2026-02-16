package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func HexToRGB(hex string) (string, error) {
	hex = strings.TrimPrefix(hex, "#")

	if len(hex) != 6 {
		return "", fmt.Errorf("invalid hex color")
	}

	r, err := strconv.ParseInt(hex[0:2], 16, 0)
	if err != nil {
		return "", err
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 0)
	if err != nil {
		return "", err
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 0)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d, %d, %d", r, g, b), nil
}
