package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("É maior que 15")
	} else {
		fmt.Println("Não é maior que 15")
	}

	if outroNum := numero; outroNum > 0 {
		fmt.Println("numero é maior do que 0")
	} else {
		fmt.Println("numer é menor do que 0")
	}

}
