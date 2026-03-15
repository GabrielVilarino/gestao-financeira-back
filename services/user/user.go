package user

import "unicode"

func validatePassword(password string) bool {
	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	// Caracteres especiais não permitidos
	forbiddenChars := map[rune]bool{
		'\'': true,
		',':  true,
		'.':  true,
		'-':  true,
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			// Verifica se é um caractere especial permitido
			if !forbiddenChars[char] {
				hasSpecial = true
			}
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}
