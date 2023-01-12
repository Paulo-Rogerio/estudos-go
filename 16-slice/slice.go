package main

import "fmt"

func main() {
	nome := make([]string, 0)
	nome = append(nome, "paulo")
	nome = append(nome, "camilla")
	fmt.Println(nome)

	idade := make(map[string]int)
	idade["Paulo"] = 40
	idade["Camilla"] = 33
	fmt.Println(idade)
	fmt.Println(idade["Paulo"])

	val, ok := idade["Lucas"]
	fmt.Println(val, ok)
}
