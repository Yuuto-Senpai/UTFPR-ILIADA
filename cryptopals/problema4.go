package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
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
	// Abrir o arquivo contendo as strings hexadecimais.
	file, err := os.Open("prob4.txt")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	// Variáveis para armazenar o melhor resultado encontrado.
	var bestText []byte
	var bestScore int
	var bestKey byte
	var bestLine string

	// Ler cada linha do arquivo e processar cada string hexadecimal.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hexString := scanner.Text()
		encrypted, err := hex.DecodeString(hexString)
		if err != nil {
			log.Fatalf("Erro ao decodificar hexadecimal: %v", err)
		}

		// Testar todas as chaves possíveis (0x00 a 0xFF).
		for key := 0; key < 256; key++ {
			decoded := singleByteXOR(encrypted, byte(key))
			score := scoreText(decoded)
			// Atualizar o melhor resultado se uma string tiver uma pontuação maior.
			if score > bestScore {
				bestScore = score
				bestText = decoded
				bestKey = byte(key)
				bestLine = hexString
			}
		}
	}

	// Verifica se ocorreu algum erro durante a leitura do arquivo.
	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	// Exibir a string cifrada, o texto decifrado e a chave utilizada.
	fmt.Printf("String cifrada: %s\n", bestLine)
	fmt.Printf("Texto decifrado: %s\n", bestText)
	fmt.Printf("Chave utilizada: 0x%x\n", bestKey)
}
