package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func xorBuffers(buffer1, buffer2 []byte) ([]byte, error) {
	if len(buffer1) != len(buffer2) {
		return nil, fmt.Errorf("os buffers têm tamanhos diferentes")
	}
	result := make([]byte, len(buffer1))

	for i := 0; i < len(buffer1); i++ {
		result[i] = buffer1[i] ^ buffer2[i]
	}

	return result, nil
}

func main() {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"

	bytes1, err := hex.DecodeString(hex1)
	if err != nil {
		log.Fatalf("Erro ao decodificar hex1: %v", err)
	}

	bytes2, err := hex.DecodeString(hex2)
	if err != nil {
		log.Fatalf("Erro ao decodificar hex2: %v", err)
	}

	xorResult, err := xorBuffers(bytes1, bytes2)
	if err != nil {
		log.Fatalf("Erro ao aplicar XOR: %v", err)
	}

	fmt.Println(hex.EncodeToString(xorResult))
}
