package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
[(3) Ransom Note - LeetCode](https://leetcode.com/problems/ransom-note/)
*/

func TestRansomNote(t *testing.T) {
	input := [][2]string{
		{"a", "b"},
		{"aa", "ab"},
		{"aa", "aab"},
	}
	expect := []bool{false, false, true}
	for i, v := range input {
		assert.Equal(t, expect[i], canConstruct(v[0], v[1]))
	}

}

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, r := range magazine {
		m[r]++
	}

	for _, r := range ransomNote {
		if v, ok := m[r]; !ok || v == 0 {
			return false
		} else {
			m[r]--
		}
	}
	return true
}
