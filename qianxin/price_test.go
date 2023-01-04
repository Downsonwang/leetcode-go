package qianxin

import (
	"fmt"
	"testing"
)

func TestMaxprofit(t *testing.T) {
	data := make([]int,0)
	data = append(data,5,11,2,5,9,6,4)
	fmt.Println(data)
	res := maxProfit(data)
	fmt.Println(res)
}
