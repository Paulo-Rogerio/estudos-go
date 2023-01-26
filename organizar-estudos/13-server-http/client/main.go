package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	// Recebo uma resposta e um error
	// Encerra a aplicação caso encontre um error
	resp, err := http.Get("http://localhost:3000/users")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 200 => Sucess
	if resp.StatusCode != 200 {
		fmt.Println("Not Success", resp.StatusCode)
		return
	}

	// Leio o que tem no Body
	// O que recebo aqui é uma estrutura em Byte, preciso transformar isso em Go
	body, err := io.ReadAll(resp.Body)

	// Variavel responsável por armazenar meu dado transformado pelo Unmarshal
	// Um aglomerado de User
	var response []User

	// o Tipo que quero receber esse dado
	// Pego o Body que ele ta lendo ( body ), e passar como um ponteiro essa variável (response )
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error retrieving data", err.Error())
	}

	fmt.Println(response)
}
