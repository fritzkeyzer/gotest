package main

import (
	"log"
	"net"
	"strings"
)

func main() {
	s := "1.2.3.4"
	log.Println(s, "out:", parseIpString(s))

	s = "1.2.3.4:5678"
	log.Println(s, "out:", parseIpString(s))

	s = "33b0:2b3a:5b98:119b:5093:d1f3:008d:4121"
	log.Println(s, "out:", parseIpString(s))

	s = "[33b0:2b3a:5b98:119b:5093:d1f3:008d:4121]:8080"
	log.Println(s, "out:", parseIpString(s))

}

// parseIpString will extract only the ip address from an ipv4 or ipv6
// address possibly containing a port number.
// eg: '1.2.3.4:5678' -> '1.2.3.4'
// eg: '[33b0:2b3a:5b98:119b:5093:d1f3:008d:4121]:8080' -> '33b0:2b3a:5b98:119b:5093:d1f3:8d:4121'
func parseIpString(ipString string) net.IP {

	colons := strings.Count(ipString, ":")

	// ipv4 without port
	if colons == 0 {
		return net.ParseIP(ipString)
	}

	// ipv4 with port
	if colons == 1 {
		ipString = strings.Split(ipString, ":")[0] // split on ':' and keep first element
		return net.ParseIP(ipString)
	}

	// ipv6 with port
	if strings.Count(ipString, "[") == 1 && strings.Count(ipString, "]") == 1 {
		i0 := strings.Index(ipString, "[") + 1
		i1 := strings.Index(ipString, "]")

		// check validity
		if i0 >= i1 {
			return nil
		}

		return net.ParseIP(ipString[i0:i1])
	}

	return net.ParseIP(ipString)
}
