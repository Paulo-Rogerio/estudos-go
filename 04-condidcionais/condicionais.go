package main

import (
	"fmt"
)

func main (){

	salario := 1300.0

	if salario < 1000.0 {
		// salario = salario - (salario * 0.08)
		salario -= (salario * 0.08)

	} else if salario <= 1200.0 {
		// salario = salario - (salario * 0.10)
		salario -= (salario * 0.10)

	} else {
		// salario = salario - (salario * 0.15)
		salario -= (salario * 0.15)		
	}

	fmt.Println("==================")
	fmt.Println("salario:", salario)


}