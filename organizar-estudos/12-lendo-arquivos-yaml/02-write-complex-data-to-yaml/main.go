package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Developer struct {
	Id         int                `yaml:"id"`
	Sector     int                `yaml:"int"`
	Salary     float32            `yaml:"salary"`
	Tasks      []string           `yaml:"tasks"`
	DailyHours []int              `yaml:"dailyHours"`
	Languages  map[string]float32 `yaml:"languages"`
}

func main() {
	d := &Developer{
		3333877,
		3,
		9000.00,
		[]string{"feature/task1", "feature/task2"},
		[]int{1, 2, 3, 4},
		map[string]float32{"Golang": 1.19, "Bash": 4.1},
	}

	content, err := yaml.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("file.yaml", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
