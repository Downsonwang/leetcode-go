package longestcommonprefix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var strs []string
	strs = append(strs, "flower", "flow", "flight")
	s := Longestcommonprefix(strs)
	fmt.Println(s)
}
