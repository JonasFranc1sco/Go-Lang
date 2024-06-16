package main

import "fmt"

//Aprendendo struct

type usuario struct {
	nome  string
	idade int8
}

func main() {
	fmt.Println("Arquivo struct")

	//Atribuindo valores as variaveis do meu struct	dentro de uma outra variavel
	var usr1 usuario
	usr1.nome = "Jonas"
	usr1.idade = 18
	fmt.Println(usr1)

	usr2 := usuario{"João", 17}
	fmt.Println(usr2)

}
