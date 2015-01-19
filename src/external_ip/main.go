package main

import (
	"external_ip/ip"
	"fmt"
)

func main() {
	ip := ip.GetActualIPAddress()
	fmt.Printf("mon adresse IP externe : %s\n", ip)
}
