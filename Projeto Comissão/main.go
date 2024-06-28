package main

import "fmt"

type mechanic struct {
	id   int8
	name string
}

type servOrder struct {
	mechanic
	car string
}

//Fazer uma função para cada opção do CRUD dos dois structs

func main() {
	fmt.Println("Welcome to the calculator of commission")
}
