package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		i := 0
		for i < 10 {
			i++
			fmt.Println("Incrementando i")
			//Função para adicionar uma pausa de um segundo a cada vez que a função é executada
			time.Sleep(time.Second)
		}
		fmt.Println(i)
	*/

	//Forma Simples de um For convencional em Golang
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j")
		time.Sleep(time.Second)
	}

	//For com cláusula Range (para iterar)

	nomes := [3]string{"Jaum", "Jonas", "Matheus"}

	for _, nomes := range nomes {
		fmt.Println(nomes)
	}

	//Iterando uma palavra com as letras
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Jonas",
		"Sobrenome": "Francisco",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
