package main //Convert hex to base64

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

// Criar uma função para converter hexadecimal para base64
func hexToBase64(hexStr string) (string, error) {
	// Decodificando string hexadecimal para bytes
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", fmt.Errorf("Erro ao decodificar hexadecimal: %v", err)
	}

	// Codificando os bytes para base64
	base64Str := base64.StdEncoding.EncodeToString(bytes)
	return base64Str, nil
}

func main() {
	// String hexadecimal fornecida pelo desafio
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	// Chamar a função para converter Hex para Base64
	base64Str, err := hexToBase64(hexStr)
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	// Exibição do resultado
	fmt.Println("Base64:", base64Str)
}

//Assinado e Revisado por Gabriel Henrique Moreira de Souza Bueno
