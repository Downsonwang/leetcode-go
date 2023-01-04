package 广度搜索

import "fmt"

type graph struct {
	vertex int
	list map[int][]int
	found bool
}
func pop(list []int)(
	int,[]int) {
	if len(list) > 0 {
		a := list[0]
		b := list[1:]
		return a,b
	}else {
		return -1,list
	}
}
func push(list []int,value int) []int {
	 result :=  append(list,value)
	 return result
}
func (g *graph) bfs (s int, t int) {
	var queue []int
	if s == t {
		return
	}
	visited := make([]bool,g.vertex+1)
	queue = append(queue,s)
	prev := make([]int,g.vertex+1)
	i := 0
	for i < len(prev) {
		prev[i] = -1
		i++
	}
	for len(queue) != 0 {
		var w int
		w,queue = pop(queue)
		for j := 0 ; j<len(g.list[w]); j++ {
			q := g.list[w][j]
			fmt.Println(q)
			if !visited[q] {
				prev[q] = w
				if q == t {
					fmt.Println(prev)
					return
				}
				visited[q] = true
				queue = append(queue,q)
			}
		}
	}

}
