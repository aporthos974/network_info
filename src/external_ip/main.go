package main

import (
	"external_ip/ip"
	"fmt"
)

func main() {
	ip := ip.GetActualIPAddress()
	fmt.Printf("content : %s", ip)
}
