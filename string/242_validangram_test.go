package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
[(3) Valid Anagram - LeetCode](https://leetcode.com/problems/valid-anagram/)
*/

func TestValidAnagram(t *testing.T) {
	input := [][2]string{
		{"anagram", "nagaram"},
		{"rat", "car"},
	}
	expect := []bool{true, false}
	for i, v := range input {
		assert.Equal(t, expect[i], isAnagram(v[0], v[1]))
	}

}

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	for _, r := range t {
		m[r]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
