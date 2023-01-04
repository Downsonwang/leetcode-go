/*
 * @Descripttion:
 * @Author:
 * @Date: 2022-09-07 14:38:22
 * @LastEditTime: 2022-09-18 10:48:08
 */
package threesum

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, 4}
	n := ThreeSum(nums)
	fmt.Println(n)
}
