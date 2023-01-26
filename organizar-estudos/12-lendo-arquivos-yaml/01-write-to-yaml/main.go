package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Person struct {
	Name string `yaml:name`
	Age  int    `yaml:age`
}

func main() {
	p := Person{"Paulo", 40}

	// Converte a Struct Pessoa para YAML
	content, err := yaml.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	// y vai armazenar um slice de byte
	fmt.Println(content)

	fmt.Println("-------")

	// String
	fmt.Println(string(content))

	err = ioutil.WriteFile("file.yaml", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
