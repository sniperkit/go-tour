package stringers

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ipAddress IPAddr) String() string {
	var ret []string
	for _, i := range ipAddress {
		ret = append(ret, fmt.Sprint(i))
	}
	return strings.Join(ret, ".")
}
