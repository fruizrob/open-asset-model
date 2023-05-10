package network_test

import (
	"encoding/json"
	"fmt"
	"net/netip"
	"testing"

	. "github.com/owasp-amass/open-asset-model/network"

	"github.com/stretchr/testify/require"
)

func TestNetblock(t *testing.T) {
	tests := []struct {
		description string
		cidr        string
		netType     string
	}{
		{
			description: "Test successful creation of Netblock with valid IPv4 CIDR",
			cidr:        "198.51.100.0/24",
			netType:     "IPv4",
		},
		{
			description: "Test successful creation of Netblock with valid IPv6 CIDR",
			cidr:        "2001:db8::/32",
			netType:     "IPv6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			prefix, _ := netip.ParsePrefix(tt.cidr)
			netblock := Netblock{Cidr: prefix, Type: tt.netType}

			require.Equal(t, tt.cidr, netblock.Cidr.String())
			require.Equal(t, tt.netType, netblock.Type)
		})
	}

	tests = []struct {
		description string
		cidr        string
		netType     string
	}{
		{
			description: "Test successful JSON serialization of Netblock with valid IPv4 CIDR",
			cidr:        "198.51.100.0/24",
			netType:     "IPv4",
		},
		{
			description: "Test successful JSON serialization of Netblock with valid IPv6 CIDR",
			cidr:        "2001:db8::/32",
			netType:     "IPv6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			prefix, _ := netip.ParsePrefix(tt.cidr)
			netblock := Netblock{Cidr: prefix, Type: tt.netType}

			jsonData, err := json.Marshal(netblock)

			require.NoError(t, err)
			require.JSONEq(t, fmt.Sprintf(`{"cidr":"%s","type":"%s"}`, tt.cidr, tt.netType), string(jsonData))
		})
	}
}
