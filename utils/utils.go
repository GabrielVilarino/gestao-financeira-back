package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func NormalizarString(s string) string {
	// Remove acentos
	t := norm.NFD.String(s)
	result := make([]rune, 0, len(t))

	for _, r := range t {
		// Remove caracteres de acentuação
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		result = append(result, r)
	}

	// Converte pra maiúsculo
	return strings.ToUpper(string(result))
}
