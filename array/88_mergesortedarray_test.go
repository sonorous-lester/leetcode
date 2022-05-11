package array

import "testing"

/*
[(3) Merge Sorted Array - LeetCode](https://leetcode.com/problems/merge-sorted-array/)
*/

func TestMergeSortedArray(t *testing.T) {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	maxLength := m + n - 1
	mi := m - 1
	ni := n - 1

	for maxLength > 0 {
		if mi <= 0 || ni <= 0 {
			break
		}

		if nums1[mi] > nums2[ni] {
			nums1[maxLength] = nums1[mi]
			mi--
		} else {
			nums1[maxLength] = nums2[ni]
			ni--
		}
		maxLength--
	}

	for ni >= 0 {
		nums1[maxLength] = nums2[ni]
		ni--
		maxLength--
	}
}
