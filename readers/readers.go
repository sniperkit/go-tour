package readers

// A Reader
type MyReader struct{}

// Reads a stream of 'A' characters into bytes.
// Returns the number of bytes read and nil.
func (MyReader) Read(bytes []byte) (n int, err error) {
	ret := bytes[:0]
	i := 0
	for ; i < len(bytes); i++ {
		ret = append(ret, 'A')
	}
	return i, nil
}
