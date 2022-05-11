package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	input := []int{3, 2, 4}
	expect := []int{1, 2}
	actual := twoSum(input, 6)
	assert.ElementsMatch(t, expect, actual)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if v, ok := m[v]; ok {
			return []int{i, v}
		}
		m[target-v] = i
	}
	return nil
}
