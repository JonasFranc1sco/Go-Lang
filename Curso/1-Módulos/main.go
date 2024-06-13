package main

import (
	"fmt"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Criando um módulo em Go")
	auxiliar.Escrever()
}

/*
go mod init modulo
Esse comando será compilado pelo comando "go build"
no terminal, que unirá todos os códigos do mesmo arquivo em um só (em binário)*/

/*Nos pacotes em Go só se pode importar funções que começam com letra maiúscula
Funções com letra minúscula só pode ser chamada dentro de seu próprio pacote*/
