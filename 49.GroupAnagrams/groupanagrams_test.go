package _9_GroupAnagrams

import "testing"

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat","tea","tan","ate","nat","bat"}
	res := groupAnagrams(strs)
	t.Log(res)
}