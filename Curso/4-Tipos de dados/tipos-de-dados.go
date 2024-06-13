package main

import (
	"errors"
	"fmt"
)

func main() {
	//int existem int8 (até 8 bits), int16(até 16 bits), int32...
	numero := 16

	fmt.Println(numero)

	// Alias
	// int32 = rune
	var numero2 rune = 1234
	fmt.Println(numero2)

	// int8 = byte
	var numero3 byte = 4
	fmt.Println(numero3)

	// Floats
	var numerofloat1 float32 = 1234.567
	fmt.Println(numerofloat1)
	// Floats só podem ser explicitos caso seja declarado o tamanho de bytes necessário (32 ou 64)

	// Float sem ser explicito
	numerofloat2 := 12345.8
	fmt.Println(numerofloat2)

	//Strings
	var str string = "Essa é uma string"
	fmt.Println(str)

	str2 := "Essa é outra string"
	fmt.Println(str2)

	/*'char' (o mais proximo de char em Go seria a aspas simples, no qual é possível verificar o valor de um
	caractere na tabela ASCII */
	char := 'b' //66 na tabela ASCII
	fmt.Println(char)

	// Dado booleano
	// True ou false
	var trueoufalse bool = true
	fmt.Println(trueoufalse)

	// Dado ERROR
	// Errors.new usado para mostrar um erro
	// Muito usado em Go
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
