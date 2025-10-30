package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func AplicarCifraDeVernan(texto string, chave string) (string, error) {
	if texto == "" || chave == "" {
		return "", errors.New("texto e chave não podem ser vazios")
	}

	// Detecta se o texto é binário (só contém 0 e 1)
	isBinary := strings.IndexFunc(texto, func(r rune) bool {
		return r != '0' && r != '1'
	}) == -1

	// Caso 1: Texto comum → cifrar
	if !isBinary {
		textoRunes := []rune(texto)
		chaveRunes := []rune(chave)

		if len(chaveRunes) < len(textoRunes) {
			return "", errors.New("a chave deve ser do mesmo tamanho ou maior que o texto")
		}

		var builder strings.Builder
		for i, c := range textoRunes {
			cifrado := c ^ chaveRunes[i]
			builder.WriteString(fmt.Sprintf("%016b", cifrado)) // escreve como binário
		}
		return builder.String(), nil
	}

	// Caso 2: Texto binário → decifrar
	if len(texto)%16 != 0 {
		return "", errors.New("texto cifrado inválido (comprimento não é múltiplo de 16)")
	}

	var builder strings.Builder
	chaveRunes := []rune(chave)

	// número de runes codificados no binário
	numRunes := len(texto) / 16
	if len(chaveRunes) < numRunes {
		return "", errors.New("a chave deve ser do mesmo tamanho ou maior que o texto original")
	}

	for i := 0; i < numRunes; i++ {
		binSegment := texto[i*16 : (i+1)*16]
		value, err := strconv.ParseInt(binSegment, 2, 32)
		if err != nil {
			return "", fmt.Errorf("erro ao converter binário: %v", err)
		}

		decifrado := rune(value) ^ chaveRunes[i]
		builder.WriteRune(decifrado)
	}

	return builder.String(), nil
}
