package LineCounter

type LCount int

func (c *LCount) Write(p []byte) (int, error) {
	ret := 0
	for i, ch := range p {
		if ch == '\n' || i == len(p)-1 {
			ret += 1
		}
	}
	*c += LCount(ret)
	return ret, nil
}
