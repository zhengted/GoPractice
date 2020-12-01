package ByteCounter

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int,error)  {
	*b += ByteCounter(len(p))
	return len(p), nil
}
