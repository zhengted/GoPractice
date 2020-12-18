package main

import "math"

func findTheDifference(s string, t string) byte {
	if len(s) <= 0 {
		return t[0]
	}
	var sNum int32
	for _, sb := range s {
		sNum += sb
	}
	var tNum int32
	for _, tb := range t {
		tNum += tb
	}
	ret := int32(math.Abs(float64(sNum - tNum)))
	return byte(ret)
}
