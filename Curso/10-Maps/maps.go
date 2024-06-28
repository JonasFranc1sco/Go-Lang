package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps!")
	//Dentro dos colchetes fica o tipo das chaves usadas no map, fora é o tipo dos valores usados no map
	usuario := map[string]string{
		//Chave       Valores
		"nome":      "Jonas",
		"Sobrenome": "Francisco",
	}
	fmt.Println(usuario)
	//Buscando específico
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Jonasu",
			"segundo":  "Franciscou",
		},

		"curso": {
			"faculdade": "Sapiens",
			"curso":     "Sistema de Informação",
		},
	}

	fmt.Println(usuario2)

	//Deletar uma chave do map
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	//Adicionar uma chave no map
	usuario2["cidade"] = map[string]string{
		"cidade": "Porto Velho",
	}

	fmt.Println(usuario2)
}
