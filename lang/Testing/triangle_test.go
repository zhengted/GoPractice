package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{12, 35, 37},
		{8, 15, 17},
	}
	for _, tt := range tests {
		if actual := calcTiangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d,%d);"+
				"got %d; Expected:%d", tt.a, tt.b, actual, tt.c)
		}
	}
}
