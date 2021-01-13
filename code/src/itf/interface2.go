package itf

type ByteCouter int

func (bc *ByteCouter) Write(p []byte) (n int, err error) {
	*bc += ByteCouter(len(p))
	return len(p), nil
}
