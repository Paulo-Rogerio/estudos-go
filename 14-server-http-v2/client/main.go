package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Monitor struct {
	Id         string `json:id`
	Monitor_Id string `json:monitor_id`
	Name       string `json:name`
}

func main() {

	// post
	jsonData := map[string]string{"monitor_id": "3", "name": "teste-1"}
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
