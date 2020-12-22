package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	s := "123456579"
	expected := []int{123, 456, 579}
	actual := splitIntoFibonacci(s)
	for i, n := range actual {
		if n != expected[i] {
			panic(errors.New(fmt.Sprintf("expected %v, got %v\n", expected, actual)))
		}
	}

}
