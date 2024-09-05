package main //Fixed XOR
import (
	"encoding/hex"
	"fmt"
	"log"
)

//Criar a Função para aplicar XOR entre duas matrizes(arranjos) de Bytes!

func fixedXOR(a, b []byte) ([]byte, error) {
	//Checar se as strings tem o mesmo Length!
	if len(a) != len(b) {
		return nil, fmt.Errorf("As strings devem possuir o mesmo comprimento!")
	}

	//Criar uma matriz para armazenar o resultado
	result := make([]byte, len(a))

	//Aplicar o XOR byte a BYTE
	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}

	return result, nil
}

func main() {

	//checar as strings hexadecimais (DONT FORGET)
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	// Converter as strings hexadecimais para bytes
	bytes1, err := hex.DecodeString(hex1)
	if err != nil {
		log.Fatalf("Erro ao decodificar hex1: %v", err)
	}

	bytes2, err := hex.DecodeString(hex2)
	if err != nil {
		log.Fatalf("Erro ao decodificar hex2: %v", err)
	}

	// Aplicar a operação XOR
	result, err := fixedXOR(bytes1, bytes2)
	if err != nil {
		log.Fatalf("Erro ao aplicar XOR: %v", err)
	}

	// Converter o resultado para hexadecimal
	resultHex := hex.EncodeToString(result)

	// Exibir o resultado final
	fmt.Println("Resultado XOR em hexadecimal:", resultHex)
}

//Assinado e Revisado por Gabriel Henrique Moreira de Souza Bueno
