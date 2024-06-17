package main

import "fmt"

// Fazendo com que o struct estudante pegue como herança as mesmas variáveis de pessoa
type pessoa struct {
	nome      string
	idade     uint8
	altura    uint8
	sobrenome string
}

// Para dar uma "herança" em Go se deve apenas colocar o nome do struct dentro do outro.
type estudante struct {
	pessoa
	curso     string
	faculdade string
	periodo   int8
}

func main() {
	fmt.Println("Herança só que não")
	//Criando uma variável preenchendo os dados do struct
	ps1 := pessoa{"Jonas", 18, 183, "Francisco"}
	fmt.Println(ps1)
	//instanciando uma variavel com os dados para o outro struct
	est1 := estudante{ps1, "Si", "Sapiens", 1}
	fmt.Println(est1)
	//Como acessar campos do struct preenchido.
	fmt.Println(est1.nome)
	fmt.Println(est1.idade)
}
