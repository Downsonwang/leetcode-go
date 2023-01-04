package main

import "fmt"

var Numbers = [...]int{
	1,10,2,6,5,
}

const  (
	target = 3
)
func main()  {
	var dataMap = make(map[int]int)
	for i,v := range Numbers {
		if p,ok :=dataMap[target-v];ok{
			fmt.Println([]int{p,i})
		}
		dataMap[v] = i
	}
}