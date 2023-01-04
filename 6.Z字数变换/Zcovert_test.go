package __Z字数变换

import (
	"fmt"
	"testing"
)

func TestZcovert(t *testing.T) {
	var s string
	s = convert("PAYPALISHIRING", 3)
	fmt.Println(s)
}
