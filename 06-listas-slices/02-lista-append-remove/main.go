package main

import (
    "fmt"
)

func RemoveCopy(slice []int, i int) []int {
    copy(slice[i:], slice[i+1:])
    return slice[:len(slice)-1]
}

func main() {
    all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

    fmt.Println("all: ", all) //[0 1 2 3 4 5 6 7 8 9]

    all = RemoveCopy(all, 5)

    fmt.Println("all: ", all) //[0 1 2 3 4 6 7 8 9]

}