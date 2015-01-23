package local

import (
	"log"
	"net"
	"regexp"
	"strings"
)

func GetLocalIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("Erreur dans la récupération de l'IP local : %s", err.Error())
	}
	for _, addr := range addrs {
		isIPv4, _ := regexp.MatchString("^\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3}/24$", addr.String())
		if isIPv4 {
			return strings.Replace(addr.String(), "/24", "", 1)
		}
	}
	return "non trouvée"
}
