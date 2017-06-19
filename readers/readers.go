package readers

type MyReader struct{}

func (MyReader) Read(bytes []byte) (n int, err error) {
	ret := bytes[:0]
	i := 0
	for ; i < len(bytes); i++ {
		ret = append(ret, 'A')
	}
	return i, nil
}
