package main

import (
	"encoding/hex"
	"fmt"
)

// Função que aplica o XOR com uma chave repetida
func repeatingKeyXOR(input, key []byte) []byte {
	output := make([]byte, len(input))

	// Aplicar XOR com a chave repetida
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i%len(key)]
	}

	return output
}

func main() {
	// Texto de entrada
	plainText := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"

	// Chave repetida
	key := "ICE"

	// Aplicar a cifra XOR com chave repetida
	cipherText := repeatingKeyXOR([]byte(plainText), []byte(key))

	// Converter o resultado para hexadecimal e exibir
	fmt.Println(hex.EncodeToString(cipherText))
}
