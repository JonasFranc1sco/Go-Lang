package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Verificador de email")
	checkmail.ValidateFormat("devbook@gmail.com")

	//Check se o Email não for um email, ele dará erro.
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)

}
