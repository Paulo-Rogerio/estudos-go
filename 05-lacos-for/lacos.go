package main

import (
	"fmt"
)

func main (){

	texto := "palavra"
	fmt.Println("Tamnho Texto:", len(texto))

	fmt.Println("=============")

    for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("=============")

	for i := 0; i < len(texto); i++ {
		fmt.Println("Index:", i)
		fmt.Println("Letra em Byte:", texto[i])
		fmt.Println("Letra em String:", string(texto[i]))
		fmt.Println("-----")
	}

	fmt.Println("=============")

	for i := 0; i < len(texto); i++ {
		if string(texto[i]) == "r" {
			break
		}
		fmt.Println(string(texto[i]))
	}

	fmt.Println("=============")

	for i := 0; i < len(texto); i++ {
		if string(texto[i]) == "r" {
			continue
		}
		fmt.Println(string(texto[i]))
	}

	fmt.Println("== Cade o While? ==")	


	tamanho := len(texto);
    i := 0

	for i < tamanho {
		fmt.Println(string(texto[i]))
		i++
	}


}