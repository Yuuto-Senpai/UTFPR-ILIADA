package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	// Converter string hexadecimal para bytes
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Erro ao decodificar string hexadecimal:", err)
		return
	}

	// Codificar bytes em base64
	base64String := base64.StdEncoding.EncodeToString(bytes)

	// Imprimir string base64 resultante
	fmt.Println(base64String)
}
