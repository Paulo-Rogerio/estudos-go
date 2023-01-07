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
	Id         string `json:id yaml:"id"`
	Monitor_Id string `json:monitor_id yaml:"monitor_id"`
	Name       string `json:name yaml:"name"`
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
		log.Println("error: %v", err)
	}

	// post
	jsonData := map[string]string{"monitor_id": f.Monitor_Id, "name": f.Name}
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
