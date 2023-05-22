package main

import (
	"log"
	"os"

	"github.com/albertodiazz/viewWifiAccesCdmx/models"
	fromCsv "github.com/albertodiazz/viewWifiAccesCdmx/tasks"
	"gopkg.in/yaml.v2"
)

// TODO
// [] Consumir los datos de forma automatica y programada
// [] Middleware para api
// [] Api
// [] Middleware para websocket
// [] Preview cliente

func ConfigFile() models.ConfigFile {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var config models.ConfigFile
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func main() {
	dat := ConfigFile()
	fromCsv.GetData(dat)

}
