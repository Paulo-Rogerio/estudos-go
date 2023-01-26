package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	dir, err := os.Open("teste/")
	if err != nil {
		panic(err)
	}

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("ERROR: %v \n", err)
		}

		fmt.Println(files[0].Name())
	}

}
