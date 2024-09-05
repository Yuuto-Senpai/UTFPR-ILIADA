package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func singleByteXOR(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i := range input {
		result[i] = input[i] ^ key
	}
	return result
}

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
	// Abrir o arquivo contendo as strings hexadecimais
	file, err := os.Open("prob4.txt")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	var bestText []byte
	var bestScore int
	var bestKey byte
	var bestLine string

	// Ler cada linha do arquivo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hexString := scanner.Text()
		encrypted, err := hex.DecodeString(hexString)
		if err != nil {
			log.Fatalf("Erro ao decodificar hexadecimal: %v", err)
		}

		// Testar todas as chaves possíveis (0x00 a 0xFF)
		for key := 0; key < 256; key++ {
			decoded := singleByteXOR(encrypted, byte(key))
			score := scoreText(decoded)
			if score > bestScore {
				bestScore = score
				bestText = decoded
				bestKey = byte(key)
				bestLine = hexString
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	// Exibir o resultado
	fmt.Printf("String cifrada: %s\n", bestLine)
	fmt.Printf("Texto decifrado: %s\n", bestText)
	fmt.Printf("Chave utilizada: 0x%x\n", bestKey)
}
