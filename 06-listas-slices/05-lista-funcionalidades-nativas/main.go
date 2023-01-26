package main

import (
    "fmt"
)

// criar lista vazia

func main() {
    lista := []int{1,2,3,4,5,6,7,8,9}

    // Pego os 3 primeiros indexes
    fmt.Println(lista[:3])
    fmt.Println(lista[:]) 

    // Pego a partir do Quarto index
    fmt.Println(lista[4:])    

    // Ultimo Item ( Tamnanho da Lista (9) - 1) = 8
    // Index é iniciado na posição 0
    fmt.Println(lista[len(lista)-1:])    
}