package main

import (
	"encoding/hex"
	"fmt"
)

// repeatingKeyXOR aplica a cifra XOR com chave repetida.
// Parâmetros:
//
//	input ([]byte): A mensagem que será cifrada.
//	key ([]byte): A chave usada para a cifra, repetida ciclicamente.
//
// Retorno:
//
//	[]byte: A sequência de bytes resultante após aplicar o XOR com a chave repetida.
func repeatingKeyXOR(input, key []byte) []byte {
	result := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = input[i] ^ key[i%len(key)]
	}
	return result
}

func main() {
	// Texto claro de entrada
	plaintext := "Burning 'em, if you ain't quick and nimble"
	// Chave que será repetida
	key := "ICE"

	// Aplicar a cifra XOR com chave repetida
	ciphertext := repeatingKeyXOR([]byte(plaintext), []byte(key))

	// Converter o resultado para hexadecimal e imprimir
	fmt.Println(hex.EncodeToString(ciphertext))
}
