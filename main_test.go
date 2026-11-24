package main

import (
	"reflect"
	"testing"
)

func TestSubsetSums(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{3, 1, 2}, []int{0, 1, 2, 3, 3, 4, 5, 6}},
		{[]int{5}, []int{0, 5}},
		{[]int{}, []int{0}},
		{[]int{-1, 1}, []int{-1, 0, 0, 1}},
	}
	for _, c := range cases {
		if got := SubsetSums(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("SubsetSums(%v) = %v, want %v", c.in, got, c.want)
		}
	}
	if n := len(SubsetSums([]int{1, 2, 3, 4, 5})); n != 32 {
		t.Errorf("expected 2^5=32 sums, got %d", n)
	}
}
