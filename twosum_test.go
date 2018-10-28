package main

import "testing"

func TestTwoSum(t *testing.T) {
	checkTwoSum(t, 3, []int{1, 2, 3}, true)
	checkTwoSum(t, 3, []int{1, 4, 3}, false)
	checkTwoSum(t, 4, []int{1, 2, 1}, false)
	checkTwoSum(t, 4, []int{2, 2, 3}, true)
	checkTwoSum(t, 4, []int{-1, 2, 3}, false)
	checkTwoSum(t, 4, []int{-1, 2, 5}, true)
	checkTwoSum(t, 14, []int{2, 12, 3}, true)
	checkTwoSum(t, 14, []int{1, 1, 2, 3, 10, 4}, true)
}

func checkTwoSum(t *testing.T, s int, a []int, e bool) {
	r := TwoSum(s, a)
	if r != e {
		t.Errorf("Expected %t from (%d, %d) but got %t", e, s, a, r)
	}
}

func TwoSum(t int, a []int) bool {
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range a {
		if v*2 == t {
			// Special case where we need to check for two instances of v
			if m[v] >= 2 {
				return true
			}
		} else if m[t-v] > 0 {
			return true
		}
	}
	return false
}
