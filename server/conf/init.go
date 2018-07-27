package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var Conf conf

type conf struct {
	Web struct {
		Port          int    `json:"port,omitempty"`
		TemplatesPath string `json:"templates_path,omitempty"`
	} `json:"web,omitempty"`
}

func init() {
	file, err := os.Open("config/server.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &Conf)
}
