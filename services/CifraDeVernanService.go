package services

import (
	"errors"
	"math/rand"
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

func Criptografar(texto string) (string, string, error) {
	textoRunes := []rune(texto)
	chave := GerarChaveAleatoria(len(textoRunes))

	textoCriptografado, err := AplicarCifraDeVernan(texto, chave)
	if err != nil {
		return "", "", err
	}

	return textoCriptografado, chave, nil
}

func GerarChaveAleatoria(tamanho int) string {
	chave := make([]rune, tamanho)

	for i := 0; i < tamanho; i++ {
		chave[i] = rune(rand.Intn(95) + 32) // ASCII imprimível: 32-126
	}

	return string(chave)
}
