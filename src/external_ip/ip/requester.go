package ip

import (
	"github.com/moovweb/gokogiri"
	"io/ioutil"
	"log"
	"net/http"
)

func GetActualIPAddress() string {
	response, _ := http.Get("http://www.adresseip.com")
	body, _ := ioutil.ReadAll(response.Body)

	document, err := gokogiri.ParseHtml(body)
	if err != nil {
		log.Fatalf("Erreur lors du parsing du document : %s", err.Error())
	}

	result, _ := document.Search("//div[@class='post']/div[@class='post-text']//b")
	return result[0].FirstChild().Content()
}
