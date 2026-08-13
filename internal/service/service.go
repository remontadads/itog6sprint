package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("empty input")
	}

	isMorse := true
	for _, r := range trimmed {
		if r != '.' && r != '-' && r != ' ' && r != '/' {
			isMorse = false
			break
		}
	}

	if isMorse {
		return morse.ToText(trimmed), nil
	}
	return morse.ToMorse(trimmed), nil
}
