package fromCsv

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/albertodiazz/viewWifiAccesCdmx/models"
)

func DownloadToCSV(url string) {
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
	DownloadToCSV(href)
}
