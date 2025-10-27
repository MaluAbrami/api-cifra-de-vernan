package services

import (
	"errors"
)

func AplicarCifraDeVernan(texto string, chave string) (string, error) {
	textoRunes := []rune(texto)
	chaveRunes := []rune(chave)

	// Valida tamanho da chave
	if len(chaveRunes) < len(textoRunes) {
		return "", errors.New("A chave deve ser do mesmo tamanho ou maior que o texto.")
	}

	result := make([]rune, len(textoRunes))

	for i, c := range textoRunes {
		result[i] = c ^ chaveRunes[i] // aplica XOR entre cada caractere
	}

	return string(result), nil
}
