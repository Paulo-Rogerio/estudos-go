package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("------")

	nomes := []string{"abolbora", "abacate", "caju", "manga"}

	for key, value := range nomes {
		fmt.Println(key, value)
	}

	fmt.Println("------")

	// Ignoro o Key
	for _, value := range nomes {
		fmt.Println(value)
	}

	fmt.Println("------")

	var i int
	for i < len(nomes) {
		fmt.Println(nomes[i])
		i++
	}

}
