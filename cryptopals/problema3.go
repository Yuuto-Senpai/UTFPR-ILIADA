package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// Função que aplica XOR byte a byte entre a mensagem cifrada e o possível valor de um byte
func singleByteXOR(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i := range input {
		result[i] = input[i] ^ key
	}
	return result
}

// Função para calcular a pontuação de legibilidade baseada na frequência de caracteres em inglês
func scoreText(text []byte) int {
	frequencyTable := "ETAOIN SHRDLUetaoinshrdlu"
	score := 0
	for _, b := range text {
		if strings.ContainsRune(frequencyTable, rune(b)) {
			score++
		}
	}
	return score
}

func main() {
	// Hex string fornecida pelo desafio
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	// Converter a string hexadecimal para bytes
	encrypted, _ := hex.DecodeString(hexString)

	// Variáveis para armazenar o melhor texto e a chave
	var bestText []byte
	var bestScore int
	var bestKey byte

	// Testar todos os valores possíveis de um byte (0x00 a 0xFF)
	for key := 0; key < 256; key++ {
		decoded := singleByteXOR(encrypted, byte(key))
		score := scoreText(decoded)
		if score > bestScore {
			bestScore = score
			bestText = decoded
			bestKey = byte(key)
		}
	}

	// Exibir o resultado mais legível e a chave usada
	fmt.Printf("Texto decifrado: %s\n", bestText)
	fmt.Printf("Chave utilizada: 0x%x\n", bestKey)
}
