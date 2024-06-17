// Ponteiros salvam endereços de memória
package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++

	//Ponteiro é uma referência de memória
	//Guarda um valor inteiro
	var var3 int = 150
	//Guarda o endereço de memória do valor inteiro
	var pont1 *int = &var3
	fmt.Println(var3, pont1)

	//Desreferenciação (Apontando para o valor dentro do endereço de memória)
	fmt.Println(var3, *pont1)
}
