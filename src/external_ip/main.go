package main

import (
	"external_ip/ip"
	"external_ip/ip/local"
	"fmt"
)

func main() {
	ip := ip.GetExternalIPAddress()
	ipLocal := local.GetLocalIPAddress()
	fmt.Printf("IP externe : %s\n", ip)
	fmt.Printf("IP local : %s\n", ipLocal)
}
