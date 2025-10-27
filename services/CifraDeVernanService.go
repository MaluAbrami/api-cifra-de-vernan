package services

func Criptografar(texto string, chave string) string {
	result := make([]rune, len(texto))

	for i, c := range texto {
		result[i] = c + rune(len(chave))
	}

	return string(result)
}

func Descriptografar(texto string, chave string) string {
	result := make([]rune, len(texto))

	for i, c := range texto {
		result[i] = c - rune(len(chave))
	}

	return string(result)
}
