package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// singleByteXOR aplica a operação XOR entre cada byte da entrada e um único valor de chave (byte).
// Parâmetros:
//
//	input ([]byte): A mensagem cifrada representada como uma sequência de bytes.
//	key (byte): O valor do byte utilizado na operação XOR para descriptografar.
//
// Retorno:
//
//	[]byte: A sequência de bytes resultante após aplicar o XOR.
func singleByteXOR(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	// Aplica XOR entre cada byte da entrada e a chave
	for i := range input {
		result[i] = input[i] ^ key
	}
	return result
}

// scoreText avalia a legibilidade do texto baseado na frequência de letras comuns em inglês.
// A função pontua textos que possuem maior presença de letras comuns na língua inglesa.
// Parâmetros:
//
//	text ([]byte): O texto que será avaliado.
//
// Retorno:
//
//	int: A pontuação, onde valores mais altos indicam uma maior probabilidade de ser um texto em inglês.
func scoreText(text []byte) int {
	// Letras comuns na língua inglesa, usadas como heurística para determinar legibilidade
	frequencyTable := "ETAOIN SHRDLUetaoinshrdlu"
	score := 0
	// Para cada byte no texto, incrementa a pontuação se for uma letra comum
	for _, b := range text {
		if strings.ContainsRune(frequencyTable, rune(b)) {
			score++
		}
	}
	return score
}

func main() {
	// Hex string fornecida pelo problema, representando a mensagem cifrada.
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	// Converte a string hexadecimal para uma sequência de bytes para aplicar o XOR.
	encrypted, _ := hex.DecodeString(hexString)

	// Variáveis para armazenar o melhor texto encontrado e a melhor chave.
	var bestText []byte
	var bestScore int
	var bestKey byte

	// Testa todas as chaves possíveis de 0x00 a 0xFF para encontrar o texto original.
	for key := 0; key < 256; key++ {
		// Aplica XOR entre a mensagem cifrada e o valor da chave.
		decoded := singleByteXOR(encrypted, byte(key))
		// Calcula a pontuação de legibilidade para o texto resultante.
		score := scoreText(decoded)
		// Atualiza o melhor resultado se o novo texto tiver uma pontuação maior.
		if score > bestScore {
			bestScore = score
			bestText = decoded
			bestKey = byte(key)
		}
	}

	// Exibe o texto mais legível e a chave utilizada para decifrá-lo.
	fmt.Printf("Texto decifrado: %s\n", bestText)
	fmt.Printf("Chave utilizada: 0x%x\n", bestKey)
}
