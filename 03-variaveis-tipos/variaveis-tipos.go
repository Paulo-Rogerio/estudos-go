package main

import "fmt"

func main (){
	a := 1
	b := "Nome"
	c := 3.14
	d := true

	fmt.Printf("Variável: %v \n", a)
	fmt.Printf("Tipo: %T \n", a)
	fmt.Println("==================")
	fmt.Printf("Variável: %v \n", b)
	fmt.Printf("Tipo: %T \n", b)
	fmt.Println("==================")
	fmt.Printf("Variável: %v \n", c)
	fmt.Printf("Tipo: %T \n", c)
	fmt.Println("==================")
	fmt.Printf("Variável: %v \n", d)
	fmt.Printf("Tipo: %T \n", d)
	fmt.Println("==================")
}