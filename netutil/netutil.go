package netutil

import (
	"fmt"
	"strings"
)

// TrimIpAddressPort will remove the port number from an ipv4 or ipv6 address string.
// eg: '1.2.3.4:5678' -> '1.2.3.4'
// eg: '[33b0:2b3a:5b98:119b:5093:d1f3:008d:4121]:8080' -> '33b0:2b3a:5b98:119b:5093:d1f3:8d:4121'
func TrimIpAddressPort(ipString string) (string, error) {
	ipString = strings.TrimSpace(ipString)

	colons := strings.Count(ipString, ":")

	// ipv4 without port
	if colons == 0 {
		return ipString, nil
	}

	// ipv4 with port
	if colons == 1 {
		ipString = strings.Split(ipString, ":")[0] // split on ':' and keep first element
		return ipString, nil
	}

	// ipv6 with port
	if strings.Count(ipString, "[") == 1 && strings.Count(ipString, "]") == 1 {
		i0 := strings.Index(ipString, "[") + 1
		i1 := strings.Index(ipString, "]")

		// check validity: if not, return original string
		if i0 >= i1 {
			return "", fmt.Errorf("invalid ip address(%s): square brackets issue", ipString)
		}

		return ipString[i0:i1], nil
	}

	// ipv6 without port
	return ipString, nil
}
