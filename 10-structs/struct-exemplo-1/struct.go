package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ClienteNome string

type Cliente struct {
	Nome  string
	Email string
	Cpf   int
}

type ClienteInternacional struct {
	Nome  string
	Email string
	Cpf   int
	Pais  string
}

type ClienteInternacionalConcatenado struct {
	Cliente
	Pais string
	// Criando uma anotação para esse tipo de atributo ( renomeando )
	Pais2 string `json:"abobrinhas"`
}

// Criando funçao que esteja atrelada a minha struct
// Struct tem um Method => Fazendo analogia a classes e methodos
// Essa variavel c que me da acesso aos atributos da struct

func (c Cliente) exibe() {
	fmt.Println("Exibindo nome do cliente pelo method: ", c.Nome)
}

func main() {

	cliente1 := Cliente{
		Nome:  "Paulo",
		Email: "p@p.com",
		Cpf:   1234,
	}

	fmt.Println(cliente1)
	fmt.Printf("Nome: %s Email: %s Cpf: %d \n", cliente1.Nome, cliente1.Email, cliente1.Cpf)
	fmt.Println("-------------")

	cliente2 := ClienteInternacional{
		Nome:  "Davi",
		Email: "d@d.com",
		Cpf:   123,
		Pais:  "Brasil",
	}

	fmt.Println(cliente2)
	fmt.Printf("Nome: %s Email: %s Cpf: %d Pais: %s \n", cliente2.Nome, cliente2.Email, cliente2.Cpf, cliente2.Pais)
	fmt.Println("-------------")

	// Composição / Heranca
	cliente3 := ClienteInternacionalConcatenado{
		Cliente: Cliente{
			Nome:  "Joao",
			Email: "j@j.com",
			Cpf:   123,
		},
		Pais:  "EUA",
		Pais2: "China",
	}

	fmt.Println(cliente3)
	fmt.Printf("Nome: %s Email: %s Cpf: %d Pais: %s PaisRenomeado: %s \n", cliente3.Cliente.Nome, cliente3.Cliente.Email, cliente3.Cliente.Cpf, cliente3.Pais, cliente3.Pais2)
	fmt.Printf("Nome: %s Email: %s Cpf: %d Pais: %s \n", cliente3.Nome, cliente3.Email, cliente3.Cpf, cliente3.Pais)
	fmt.Println("-------------")

	cliente1.exibe()
	cliente3.exibe()

	// Transformando algo de uma struct em json
	clienteJson, err := json.Marshal(&cliente3)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Isso vai retorna um slice de bytes por conta do Marshal, precisamos converter
	// Resposta [123 125]
	fmt.Println(clienteJson)

	// Convertendo para string
	// {"Nome":"Joao","Email":"j@j.com","Cpf":123,"Pais":"EUA","abobrinhas":"China"}
	fmt.Println(string(clienteJson))

	fmt.Println("-------------")

	// Hidratando / Preenchendo minha struct

	jsoncliente4 := `{"Nome":"Maria","Email":"m@m.com","Cpf":444,"Pais":"Brasil","abobrinhas":"Ceara"}`
	cliente4 := ClienteInternacionalConcatenado{}

	json.Unmarshal([]byte(jsoncliente4), cliente4)
	json.Unmarshal([]byte(jsoncliente4), &cliente4)

	fmt.Printf("Nome: %s Email: %s Cpf: %d Pais: %s PaisRenomeado: %s \n", cliente4.Cliente.Nome, cliente4.Cliente.Email, cliente4.Cliente.Cpf, cliente4.Pais, cliente4.Pais2)
	fmt.Println(cliente4)

}
