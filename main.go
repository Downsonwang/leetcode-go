package main

import (
	"fmt"
	"math"
)

func Bread(f func(city string) []string, workList []string) {
	seen := make(map[string]bool)
	for len(workList) > 0 {
		items := workList
		workList = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				workList = append(workList, f(item)...)
			}
		}

	}

}

func main() {
	num := 58
	numFloat64 := math.Pow10(num)

	fmt.Println(numFloat64)

}
