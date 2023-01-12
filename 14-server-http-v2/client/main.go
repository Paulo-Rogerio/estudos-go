package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Monitor struct {
	Id          string `json:"id" yaml:"id"`
	Template_Id string `json:"template_id" yaml:"template_id"`
	Profile     string `json:"profile" yaml:"profile"`
}

func main() {

	// Extract Yaml
	f := &Monitor{}
	source, err := ioutil.ReadFile("manifesto-gerado-meeseeks.yaml")
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal([]byte(source), &f)
	if err != nil {
		log.Println("error: ", err)
	}

	// post
	jsonData := map[string]string{"template_id": f.Template_Id, "profile": f.Profile}
	jsonValue, _ := json.Marshal(jsonData)
	resp, err := http.Post("http://localhost:3000/api/monitors", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("Error : %s\n", err)
		return
	}
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

	fmt.Println("----------")

	// get
	resp, err = http.Get("http://localhost:3000/api/monitors")
	if err != nil {
		fmt.Printf("Error : %s\n", err)
		return
	}

	data, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

}
