package ip

import (
	"github.com/moovweb/gokogiri"
	"io/ioutil"
	"log"
	"net/http"
)

func GetExternalIPAddress() string {
	body := sendRequest("http://www.adresseip.com")

	document, err := gokogiri.ParseHtml(body)
	if err != nil {
		log.Fatalf("Erreur lors du parsing du document : %s", err.Error())
	}

	result, _ := document.Search("//div[@class='post']/div[@class='post-text']//b")
	return result[0].FirstChild().Content()
}

func sendRequest(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erreur sur l'URL %s : %s", url, err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Erreur dans le résultat de l'URL %s : %s", url, err.Error())
	}
	return body
}

func searchIP(htmlContent []byte) string {
	document, err := gokogiri.ParseHtml(htmlContent)
	if err != nil {
		log.Fatalf("Erreur lors du parsing du document : %s", err.Error())
	}

	result, err := document.Search("//div[@class='post']/div[@class='post-text']//b")
	if err != nil {
		log.Fatalf("Erreur lors de la récupération de l'IP : %s", err.Error())
	}
	return result[0].FirstChild().Content()
}
