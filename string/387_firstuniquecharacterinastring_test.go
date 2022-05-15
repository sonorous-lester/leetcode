package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
[(3) First Unique Character in a String - LeetCode](https://leetcode.com/problems/first-unique-character-in-a-string/)
*/

func TestFirstUniqueCharacterInAString(t *testing.T) {
	s := []string{"leetcode", "loveleetcode", "aabb"}
	expect := []int{0, 2, -1}
	for i, v := range s {
		assert.Equal(t, expect[i], firstUniqChar(v))
	}
}

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, r := range s {
		if _, ok := m[r]; ok {
			m[r]++
		} else {
			m[r] = 1
		}
	}
	for i, r := range s {
		if v, ok := m[r]; ok && v == 1 {
			return i
		}
	}
	return -1
}
