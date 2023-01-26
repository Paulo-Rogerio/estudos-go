package main

import (
	"fmt"
	"strconv"
)

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

	var dd int = 127
	var ddInt8 int8 
	ddInt8 = int8(dd)
	fmt.Println("Conversao Numero", ddInt8)

	var ee float32 = 127.99
	var ddInt int 
	ddInt = int(ee)
	fmt.Println("Conversao Numero", ddInt)

	fmt.Println("==================")
	fmt.Println("Usando strconv")
   
	// Convertendo String True para Bool
	ff, _ := strconv.ParseBool("true")
	fmt.Printf("Valor da variavel e: %v \n", ff)
	fmt.Printf("O Tipo da variavel e: %T \n", ff)

	hh := "54.66"
    ii, _ := strconv.ParseFloat(hh, 64)
	fmt.Printf("Valor da variavel e: %v \n", ii)
	fmt.Printf("O Tipo da variavel e: %T \n", ii)

	fmt.Println("==================")
	fmt.Println("https://pkg.go.dev/strconv")






}