// Es un proyecto el cual pretende darte la ubicacion de los puntos de acceso
// de wifi de la cdmx mas cercanos. Para ello se obtiene la geolocalizacion del usuario
// y en base a los datos, decirle cual es sus opciones mas cercanas
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/albertodiazz/viewWifiAccesCdmx/models"
	fromCsv "github.com/albertodiazz/viewWifiAccesCdmx/tasks"
	"gopkg.in/yaml.v2"
)

// TODO
// [X] Consumir los datos de forma automatica y programada
// [X] Levanta servidor para consumo de json en base al csv
// 	- localhost:8080/data GET
// [X] Setup con Redis

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
	// dat := ConfigFile()

	// fromCsv.GetData(dat)
	// models.ConnectRedis(dat)

	http.HandleFunc("/data", fromCsv.ReadCsvData)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
