package main
import (
	"fmt"
	"strconv"
)

type Flags uint

func convertToBin(n int) string {
	const (
		FlagUp Flags = 1 << iota
		FlagBroadcast
		FlagLoopback
		FlagPointToPoint
		FlagMulticast
	)
	fmt.Println(FlagUp,FlagBroadcast,FlagLoopback,FlagPointToPoint,FlagMulticast)
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main(){
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(1024),
	)
}
