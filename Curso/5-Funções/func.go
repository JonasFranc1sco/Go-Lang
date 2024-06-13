package main

import "fmt"

// Criando uma função de somar dois números inteiros
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosmath(n3, n4 int8) (int8, int8) {
	soma := n3 + n4
	subtra := n3 - n4
	return soma, subtra
}

func main() {
	//Colocando uma variavel indicando dois números que serão jogados na função para a efetuação da soma
	soma := somar(10, 20)
	fmt.Println(soma)

	//Criando uma função dentro de uma variavel
	var f = func(txt string) string {
		fmt.Print(txt)
		return txt
	}
	resultado := f("Essa é a variável F")
	fmt.Println(resultado)

	// As funções em Go podem ter mais de um retorno
	// Terá de ser feita uma variável para cada valor retornável de uma Função
	resultadoSoma, resultadoSub := calculosmath(10, 15)
	fmt.Println(resultadoSoma, resultadoSub)

}
