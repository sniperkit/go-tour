package rot13

import (
	"io"
	"unicode"
)

type Reader struct {
	R io.Reader
}

func (reader Reader) Read(p []byte) (n int, err error) {
	n, err = reader.R.Read(p)
	b := p[:0]
	for i := 0; i < n; i++ {
		b = append(b, toRot13(p[i]))
	}
	return
}

func toRot13(b byte) byte {
	transform := func(b, base byte) byte {
		return base + (b-base+13)%26
	}

	switch {
	case unicode.IsLower(rune(b)):
		return transform(b, 'a')
	case unicode.IsUpper(rune(b)):
		return transform(b, 'A')
	default:
		return b
	}
}
