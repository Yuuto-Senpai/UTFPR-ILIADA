package main //singlebyte to xor

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

// Função que pontua uma string com base na frequência de caracteres em inglês
func scoreText(text string) int {
	frequency := map[rune]int{
		'e': 13, 't': 9, 'a': 8, 'o': 8, 'i': 7, 'n': 7, 's': 6, 'h': 6, 'r': 6, 'd': 4, 'l': 4, 'u': 3,
		' ': 15, '.': 2, ',': 2, '\'': 1, 'c': 2, 'g': 2, 'm': 3, 'y': 3,
	}

	score := 0
	for _, char := range strings.ToLower(text) {
		if val, exists := frequency[char]; exists {
			score += val
		}
	}
	return score
}

// Função que aplica XOR de um único byte a uma sequência de bytes
func singleByteXOR(input []byte, key byte) []byte {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key
	}
	return output
}

// Função principal para decifrar a string
func main() {
	// String hexadecimal fornecida no desafio
	hexStr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	// Converter a string hexadecimal para bytes
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		log.Fatalf("Erro ao decodificar hexadecimal: %v", err)
	}

	var bestScore int
	var bestResult string
	var bestKey byte

	// Testar todas as possíveis chaves de 0x00 a 0xFF
	for key := 0; key < 256; key++ {
		// Aplicar XOR com a chave atual
		decoded := singleByteXOR(bytes, byte(key))
		decodedText := string(decoded)

		// Calcular a pontuação baseada na frequência de caracteres
		score := scoreText(decodedText)

		// Atualizar o melhor resultado com base na pontuação
		if score > bestScore {
			bestScore = score
			bestResult = decodedText
			bestKey = byte(key)
		}
	}

	// Exibir o resultado
	fmt.Printf("Chave encontrada: 0x%x\n", bestKey)
	fmt.Println("Texto decifrado:", bestResult)
}

//Assinado e Revisado por Gabriel Henrique Moreira de Souza Bueno
