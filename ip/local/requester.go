package local

import (
	"bytes"
	"log"
	"net"
	"regexp"
	"strings"
)

func GetIPv4() string {
	return fetchIP("24")
}

func GetIPv6() string {
	return fetchIP("64")
}

func fetchIP(base string) string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("Erreur lors la récupération de l'IP local : %s", err.Error())
	}
	regex, baseSuffix := composeBaseSuffix(base)
	for _, addr := range addrs {
		if isIP(regex, addr) {
			return strings.Replace(addr.String(), baseSuffix, "", 1)
		}
	}
	return "non trouvée"
}

func isIP(baseRegexp string, addr net.Addr) bool {
	matchedIP, _ := regexp.MatchString(baseRegexp, addr.String())
	return matchedIP
}

func composeBaseSuffix(base string) (regex string, baseSuffix string) {
	bufferRegex := bytes.NewBufferString("/")
	bufferRegex.WriteString(base)
	baseSuffix = bufferRegex.String()
	bufferRegex.WriteString("$")
	return bufferRegex.String(), baseSuffix
}
