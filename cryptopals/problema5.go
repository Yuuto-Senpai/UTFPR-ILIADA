package main

import (
	"encoding/hex"
	"fmt"
)

// repeatingKeyXOR aplica a cifra XOR entre cada byte da entrada e a chave que se repete ciclicamente.
// Parâmetros:
//
//	input ([]byte): A mensagem em texto claro que será cifrada.
//	key ([]byte): A chave utilizada para a cifra. Essa chave será repetida ao longo da mensagem.
//
// Retorno:
//
//	[]byte: A sequência de bytes resultante após aplicar o XOR com a chave repetida.
func repeatingKeyXOR(input, key []byte) []byte {
	result := make([]byte, len(input))
	// Para cada byte da entrada, aplica XOR com a chave correspondente. A chave se repete ciclicamente.
	for i := 0; i < len(input); i++ {
		result[i] = input[i] ^ key[i%len(key)]
	}
	return result
}

func main() {
	// Texto claro de entrada que será cifrado.
	plaintext := "Burning 'em, if you ain't quick and nimble"
	// Chave utilizada na cifra. Ela será repetida ao longo do texto.
	key := "ICE"

	// Aplicar a cifra XOR com chave repetida sobre o texto claro.
	ciphertext := repeatingKeyXOR([]byte(plaintext), []byte(key))

	// Converter o resultado da cifra para hexadecimal e exibir o texto cifrado.
	fmt.Println(hex.EncodeToString(ciphertext))
}
