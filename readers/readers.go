package readers

// MyReader is a reader
type MyReader struct{}

// Read reads a stream of 'A' characters into bytes.
// Returns the number of bytes read and nil.
func (MyReader) Read(bytes []byte) (n int, err error) {
	ret := bytes[:0]
	i := 0
	for ; i < len(bytes); i++ {
		ret = append(ret, 'A')
	}
	return i, nil
}
