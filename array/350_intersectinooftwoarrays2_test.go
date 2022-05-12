package array

import (
	"testing"
)

/*
[(3) Intersection of Two Arrays II - LeetCode](https://leetcode.com/problems/intersection-of-two-arrays-ii/)

*/

func TestIntersectionOfTwoArrays2(t *testing.T) {
	
}

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, num := range nums1 {
		if _, ok := m[num]; ok {
			m[num]++
		} else {
			m[num] = 1
		}
	}

	ans := []int{}
	for _, num := range nums2 {
		if v, ok := m[num]; ok && v > 0 {
			ans = append(ans, num)
			m[num]--
		}
	}
	return ans
}
