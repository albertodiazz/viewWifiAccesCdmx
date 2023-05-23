package fromCsv

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/albertodiazz/viewWifiAccesCdmx/models"
)

func downloadToCSV(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error al realizar la solicitud GET: %s\n", err)
	}
	defer resp.Body.Close()

	file, err := os.OpenFile("./data/wifiCdmx.csv",
		os.O_WRONLY|os.O_CREATE,
		0644)
	if err != nil {
		log.Fatalf("Error al crear el archivo csv: %s", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalf("Error al salvar el archivo csv: %s", err)
	}
	log.Println("Archvio CSV descargado exitosamente")
}

func GetData(data models.ConfigFile) {
	resp, err := http.Get(data.Server.URLCDMX)
	if err != nil {
		log.Fatalf("Error al realizar la solicitud GET: %s\n", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Error al obtner el body: %s\n", err)
	}

	link := doc.Find(data.Server.CLASSNAME)
	href, exists := link.Attr("href")
	if !exists {
		log.Println("Url de csv para descargar no encontrada")
	}
	// Retorna el link del CSV encontrado en la pagina del gobierno
	// Esto lo hago, ya que su link lo cambian en base la fecha de upload
	// de sus datos
	log.Printf("Url del CSV: %s", href)

	/////////////////////
	// Aqui descargamos el CSV
	/////////////////////
	downloadToCSV(href)
}

func ReadCsvData(w http.ResponseWriter, r *http.Request) {
	// func ReadCsvData() {
	file, err := os.Open("./data/wifiCdmx.csv")
	if err != nil {
		log.Fatalf("No se logro abrir el documento csv: %s", err)
	}
	defer file.Close()

	// Leer el archivo CSV utilizando el paquete encoding/csv
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("No se logro leer el documento csv: %s", err)
	}

	var dw models.DatosWifi
	var data []models.DatosWifi

	for _, each := range records[1:] {
		dw.Id = each[0]
		dw.Latitud = each[3]
		dw.Longitud = each[4]
		dw.Colonia = each[5]
		dw.Alcaldia = each[6]
		data = append(data, dw)
	}

	// // Codificar los datos como JSON y enviar la respuesta al cliente
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
