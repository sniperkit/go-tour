package stringers

import (
	"fmt"
	"strings"
)

// Represents an IPv4 address
type IPAddr [4]byte

// Stringifies an IPv4 address with a dot between each octet.
func (ipAddress IPAddr) String() string {
	var ret []string
	for _, i := range ipAddress {
		ret = append(ret, fmt.Sprint(i))
	}
	return strings.Join(ret, ".")
}
