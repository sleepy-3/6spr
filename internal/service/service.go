package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(t string) (string, error) {

	l := len(t)

	if l == 0 {
		return "", errors.New("empty file")
	}
	if strings.ContainsAny(t, ".-") {
		return morse.ToText(t), nil
	} else {
		return morse.ToMorse(t), nil
	}
}
