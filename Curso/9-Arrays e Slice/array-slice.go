package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays")
	//Array explicitamente criado
	var arr1 [5]int
	arr1[0] = 0
	fmt.Println(arr1)

	//Array implicitamente criado
	arr2 := [5]int{0, 1, 2, 3, 4}
	fmt.Println(arr2)
	//Array não é tão flexível, pois não é possivel mudar diretamente um valor dentro dele
	//Criando um array sem limite de qnt de dados.
	arr3 := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(arr3)

	//Slice
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arr3))

	//Função append para acrescentar um valor no slice.
	slice = append(slice, 18)
	fmt.Println(slice)

	//Pegando uma fatia (slice) de um array [Primeiro dado inclusivo, segundo exclusivo]
	slice2 := arr3[1:4]
	fmt.Println(slice2)
	//No slice é possível mudar um valor dentro dele sem que cause erros
	arr3[2] = 78
	fmt.Println(slice2)

	//Arrays internos
	//Make serve para especificar o tamanho do slice
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Length
	fmt.Println(cap(slice3)) //Capacity

	//Adicionar um valor no slice
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	slice3 = append(slice3, 7)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]int8, 1)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
