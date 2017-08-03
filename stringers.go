package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	// At this point you could simply
	//return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])

	var string string

	for i := 0; i < len(ip); i++ {
		string += fmt.Sprintf("%v", ip[i])
		if i < len(ip) - 1 {
			string += "."
		}
	}

	return string
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
