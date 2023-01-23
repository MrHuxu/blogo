package conf

import (
	"encoding/json"
	"log"
	"os"

	"github.com/MrHuxu/blogo/config"
)

// Conf exports an instance of conf
var Conf config.Conf

func init() {
	bytes, err := os.ReadFile("config/server.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bytes, &Conf)
}
