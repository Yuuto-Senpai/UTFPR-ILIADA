package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

// hexToBase64 é uma função responsável por realizar a conversão de uma string
// hexadecimal para uma representação codificada em Base64. O algoritmo é implementado
// em duas etapas principais: primeiro, a string hexadecimal é convertida em uma sequência
// de bytes, e em seguida, essa sequência é codificada no formato Base64.
//
// Parâmetros:
//
//	hexStr (string): uma string representando valores em formato hexadecimal.
//
// Retorno:
//
//	string: a representação em Base64 da sequência de bytes derivada do valor hexadecimal original.
//	error: em caso de falha na conversão, um erro é retornado para indicar a natureza do problema.
func hexToBase64(hexStr string) (string, error) {
	// Decodificação da string hexadecimal para uma sequência de bytes.
	// Esta etapa é fundamental para a posterior codificação em Base64.
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		// Em caso de erro na decodificação, a função retorna o erro apropriado
		// e uma string vazia, permitindo o tratamento adequado pelo chamador.
		return "", err
	}

	// Codificação da sequência de bytes em formato Base64.
	// O uso de Base64 é amplamente utilizado em sistemas que necessitam
	// de uma representação textual compacta de dados binários.
	base64Str := base64.StdEncoding.EncodeToString(bytes)

	// Retorna a string codificada em Base64 e nil para indicar que a operação foi bem-sucedida.
	return base64Str, nil
}

func main() {
	// String hexadecimal de entrada. Esta string será utilizada para demonstrar a
	// conversão do formato hexadecimal para Base64. O valor escolhido representa
	// uma frase textual específica quando interpretada em ASCII.
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	// Chamada à função hexToBase64. Aqui, o programa realiza a conversão da string
	// hexadecimal para Base64. A função retorna tanto o valor convertido quanto
	// um possível erro, que deve ser tratado.
	base64String, err := hexToBase64(hexString)
	if err != nil {
		// Caso ocorra um erro durante a conversão, o programa será finalizado
		// e uma mensagem de erro será exibida no console. A utilização de log.Fatalf
		// permite uma finalização imediata e o fornecimento de uma mensagem detalhada.
		log.Fatalf("Erro ao converter hexadecimal para Base64: %v", err)
	}

	// Exibição do resultado final no formato Base64. O valor impresso representa
	// a versão codificada em Base64 da sequência de bytes derivada da string hexadecimal original.
	fmt.Println("Resultado em Base64:", base64String)
}
