package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
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

func main() {
	// Abrir o arquivo com as linhas hexadecimais
	file, err := os.Open("4.txt")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	var bestScore int
	var bestResult string
	var bestLine string
	var bestKey byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hexStr := scanner.Text()

		// Converter a string hexadecimal para bytes
		bytes, err := hex.DecodeString(hexStr)
		if err != nil {
			log.Fatalf("Erro ao decodificar hexadecimal: %v", err)
		}

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
				bestLine = hexStr
				bestKey = byte(key)
			}
		}
	}

	// Exibir o resultado final
	fmt.Printf("Linha hex codificada: %s\n", bestLine)
	fmt.Printf("Chave encontrada: 0x%x\n", bestKey)
	fmt.Println("Texto decifrado:", bestResult)
}
