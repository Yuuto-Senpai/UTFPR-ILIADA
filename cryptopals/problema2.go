package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

// xorBuffers aplica a operação XOR entre dois buffers de bytes e retorna o resultado como um slice de bytes.
// Parâmetros:
//
//	buffer1 ([]byte): Primeiro buffer de bytes para a operação XOR.
//	buffer2 ([]byte): Segundo buffer de bytes para a operação XOR.
//
// Retorno:
//
//	[]byte: Slice de bytes resultante da aplicação da operação XOR entre buffer1 e buffer2.
//	error: Retorna um erro caso os buffers tenham comprimentos diferentes.
func xorBuffers(buffer1, buffer2 []byte) ([]byte, error) {
	// Verifica se os buffers têm o mesmo comprimento, pois a operação XOR requer buffers de mesmo tamanho.
	if len(buffer1) != len(buffer2) {
		return nil, fmt.Errorf("os buffers têm tamanhos diferentes")
	}

	// Cria um slice para armazenar o resultado da operação XOR.
	result := make([]byte, len(buffer1))

	// Aplica a operação XOR byte a byte.
	for i := 0; i < len(buffer1); i++ {
		result[i] = buffer1[i] ^ buffer2[i]
	}

	// Retorna o resultado como um slice de bytes.
	return result, nil
}

func main() {
	// Definindo as strings hexadecimais de entrada. Essas strings representam os dois buffers
	// que serão utilizados para a operação XOR.
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"

	// Converte as strings hexadecimais para bytes utilizando a função hex.DecodeString.
	bytes1, err := hex.DecodeString(hex1)
	if err != nil {
		log.Fatalf("Erro ao decodificar hex1: %v", err)
	}

	bytes2, err := hex.DecodeString(hex2)
	if err != nil {
		log.Fatalf("Erro ao decodificar hex2: %v", err)
	}

	// Aplica a operação XOR entre os dois buffers de bytes e captura o resultado.
	xorResult, err := xorBuffers(bytes1, bytes2)
	if err != nil {
		log.Fatalf("Erro ao aplicar XOR: %v", err)
	}

	// Converte o resultado XOR para hexadecimal e imprime o valor final no console.
	fmt.Println("Resultado XOR em Hexadecimal:", hex.EncodeToString(xorResult))
}
