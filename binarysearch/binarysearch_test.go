package binarysearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.

*/

func TestBinarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expect := 4
	ans := binarySearch(nums, target)
	assert.Equal(t, expect, ans)
}

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	ans := -1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			ans = mid
			return ans
		}
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}
