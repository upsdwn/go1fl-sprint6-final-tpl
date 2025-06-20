package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertMorse(text string) string {
	isMorse := !strings.ContainsFunc(text, func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	})

	if isMorse {
		return morse.ToMorse(text)
	} else {
		return morse.ToText(text)
	}
}
