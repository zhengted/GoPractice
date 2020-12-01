package WordCounter

import (
	"bufio"
	"strings"
)

type Count int

func (c *Count) Write(p []byte) (int,error) {
	start, end := -1,-1
	ret := 0
	for i,ch := range p {
		if ch != ' ' && ch != 0 && start == -1 {
			start,end = i,i
		}else if ch != ' ' && i != len(p)-1 && start != -1 {
			end++
		}else {
			ret++
			start = -1
			end = -1
		}
	}
	*c += Count(ret)
	return ret,nil
}

// 更好的写法
type WordCounter int
func (w *WordCounter)  Write(p []byte) (int,error) {
	s := strings.NewReader(string(p))
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)
	sum := 0
	for bs.Scan() {
		sum++
	}
	*w = WordCounter(sum)
	return sum,nil
}
