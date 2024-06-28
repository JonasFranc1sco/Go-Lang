package main

import "fmt"

// Em Go não é necessário utilizar break para cada case.
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

// Uma outra forma de utilizar switch (mais usada para switchs com mais de uma variável como opção)
func diaDaSemana2(numero int) string {
	var diaDaSemana5 string
	switch {
	case numero == 1:
		diaDaSemana5 = "Domingo"
	case numero == 2:
		diaDaSemana5 = "Segunda"
	case numero == 3:
		diaDaSemana5 = "Terça"
	case numero == 4:
		diaDaSemana5 = "Quarta"
	case numero == 5:
		diaDaSemana5 = "Quinta"
	case numero == 6:
		diaDaSemana5 = "Sexta"
	case numero == 7:
		diaDaSemana5 = "Sábado"
	default:
		diaDaSemana5 = "Número inválido"
	}
	return diaDaSemana5
}

func main() {
	fmt.Println("Switchs")
}
