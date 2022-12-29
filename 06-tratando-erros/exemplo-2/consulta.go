package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://xpgoogle.xxcom.br")

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(resp.Status)
}
