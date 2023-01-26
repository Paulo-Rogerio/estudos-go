package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)


type Engineer struct {
	Id			int					`yaml:"id"`
	Sector  	int 				`yaml:sector`
	Nome    	string				`yaml:nome`
	Tasks		[]string			`yaml:tasks`
	Daily		[]int				`yaml:daily`
	Languages	map[string]float32	`yaml:languages`
}

func main() {
	f := &Engineer{}
	source, err := ioutil.ReadFile("file.yaml")
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal([]byte(source), &f)
	if err != nil {
		log.Println("error: %v", err)
	}

	fmt.Printf("O Valor de Id e: %d \n", f.Id)
	fmt.Printf("O Valor de Setor e: %d \n", f.Sector)
	fmt.Printf("O Valor de Nome e: %s \n", f.Nome)
	fmt.Printf("O Valor de Tasks e: %s \n", f.Tasks)
	fmt.Printf("O Valor de Daily e: %d \n", f.Daily)
	fmt.Println("O Valor de Languages e: ", f.Languages)

}