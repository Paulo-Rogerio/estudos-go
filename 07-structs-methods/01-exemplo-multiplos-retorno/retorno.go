package main

import "fmt"

func main() {
	resp, str := soma(1, 1)
	fmt.Println(resp, str)
}

// Funcao retorna 2 tipos ( int e string )
func soma(x, y int) (int, string) {
	return (x + y), "Somou"
}
