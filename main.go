package main

import (
	"fmt"
	"network_info/ip/extern"
	"network_info/ip/local"
)

func main() {
	externalIp := extern.GetExternalIPAddress()
	localIPv4 := local.GetIPv4()
	localIPv6 := local.GetIPv6()
	fmt.Printf("IP externe : %s\n", externalIp)
	fmt.Printf("IPv4 local : %s\n", localIPv4)
	fmt.Printf("IPv6 local : %s\n", localIPv6)
}
