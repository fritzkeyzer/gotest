package netutil

import (
	"reflect"
	"testing"
)

func TestTrimIpAddressPort(t *testing.T) {
	tests := []struct {
		ipString string
		want     string
	}{
		// ipv4 tests
		{
			ipString: "1.2.3.4",
			want:     "1.2.3.4",
		},
		{
			ipString: "1.2.3.4:1337",
			want:     "1.2.3.4",
		},
		{
			ipString: "255.255.255.255",
			want:     "255.255.255.255",
		},
		{
			ipString: "255.255.255.255:65000",
			want:     "255.255.255.255",
		},
		// ipv6 tests
		{
			ipString: "db1f:e739:2066:8f3a:635e:c914:cb8c:d450",
			want:     "db1f:e739:2066:8f3a:635e:c914:cb8c:d450",
		},
		{
			ipString: "[7229:2635:0b74:a2db:48dd:bace:59d9:8aee]:1337",
			want:     "7229:2635:0b74:a2db:48dd:bace:59d9:8aee",
		},
		// ipv6 strange cases
		{
			ipString: "::1",
			want:     "::1",
		},
		{
			ipString: "[::1]:8080",
			want:     "::1",
		},
	}

	for _, test := range tests {
		have, err := TrimIpAddressPort(test.ipString)
		if err != nil {
			t.Errorf("TrimIpAddressPort(%v): error: %v", test.ipString, err)
			continue
		}

		if !reflect.DeepEqual(have, test.want) {
			t.Errorf("TrimIpAddressPort(%v) = %v, want %v", test.ipString, have, test.want)
		}
	}
}
