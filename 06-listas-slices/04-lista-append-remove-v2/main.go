package main

import (
    "fmt"
)

// criar lista vazia

func main() {
    lista := []int{1,2,3,4,5,6,7,8,9}

    lista_menos_5 := make([]int, 0)

    fmt.Println(lista)
    fmt.Println(lista_menos_5)
    fmt.Println("=============")

    for i := 0; i < len(lista); i++ {

        if lista[i] != 5 {
            lista_menos_5 = append(lista_menos_5, lista[i])
        }
    }

    fmt.Println(lista)
    fmt.Println(lista_menos_5)
}