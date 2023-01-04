package _1_simplePath

import "strings"

// 71.简化路径 - 栈
func simplifyPath(path string) string {

    var stack []string
	arr := strings.Split(path,"/")
	for _,v := range arr {
		if v == ".." {
			//.. 就意味着栈顶出栈
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}else if v != "" && v != "."{
			// 进栈
			 stack = append(stack,v)
		}
	}
	return "/" + strings.Join(stack,"/")
}
