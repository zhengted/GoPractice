package main

import "testing"

func TestSubStr(t *testing.T)  {
	testData := []struct{
		str string
		res int
	}{
		// Normal Case
		{"abcabcbb",3},
		{"pwwkew",3},

		// Edge Cases
		{"",0},
		{"bbbb",1},
		{"b",1},
		{"abcabcabcd",4},

		// Chinese support
		{"一二三二一",3},
	}
	for _,tt := range testData {
		actual := lengthOfLongestSubstring(tt.str)
		if actual != tt.res {
			t.Errorf("lengthOfLongestSubstring %s expected:%d actual:%d",
				tt.str,tt.res,actual)
		}
	}
	
}

func BenchmarkSubstr(b *testing.B)  {
	s := "abcabcbb"
	for i:=0; i < 13; i++ {
		s = s + s
	}
	ans := 3

	b.Logf("len(s):%d",len(s))
	b.ResetTimer()

	for i:=0; i < b.N; i++ {
		actual := lengthOfLongestSubstring(s)
		if actual != ans {
			b.Errorf("lengthOfLongestSubstring %s expected:%d actual:%d",
				s,ans,actual)
		}
	}
}