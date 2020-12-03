package ipconv

import (
	"fmt"
	"unsafe"
)

type ipv4Error struct {
	ip string
}

func (e ipv4Error) Error() string {
	return fmt.Sprintf("ipconv: invalid ipv4 %s", e.ip)
}

// V42Long converts an ipv4 string to an uint32 integer.
// Panic if format of ip is not ipv4
func V42Long(ip string) (long uint32) {
	l := len(ip)
	if l > 15 {
		panic(ipv4Error{ip})
	}
	var (
		n uint32
		b = 24
	)
	for i := 0; i < l; i++ {
		c := ip[i]
		switch {
		case c == '.':
			if b <= 0 {
				panic(ipv4Error{ip})
			}
			long |= n << b
			n, b = 0, b-8
		case c >= '0' && c <= '9':
			n = n*10 + uint32(c-'0')
			if n > 255 {
				panic(ipv4Error{ip})
			}
		default:
			panic(ipv4Error{ip})
		}
	}

	return long | n
}

// Long2V4 convert an uint32 integer to ipv4 string
func Long2V4(ip uint32) string {
	b := make([]byte, 0, 15)

	b = appendByte(b, byte(ip>>24))
	b = append(b, '.')
	b = appendByte(b, byte(ip>>16))
	b = append(b, '.')
	b = appendByte(b, byte(ip>>8))
	b = append(b, '.')
	b = appendByte(b, byte(ip))

	return b2s(b)
}

func appendByte(dst []byte, n byte) []byte {
	switch {
	case n < 10:
		return append(dst, n+'0')
	case n < 100:
		dst = append(dst, n/10+'0')
		n -= n / 10 * 10
		return append(dst, n+'0')
	default:
		dst = append(dst, n/100+'0')
		n -= n / 100 * 100
		dst = append(dst, (n/10)+'0')
		n -= n / 10 * 10
		return append(dst, n+'0')
	}
}

// b2s converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func b2s(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}
