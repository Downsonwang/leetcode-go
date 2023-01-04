/*
 * @Descripttion:
 * @Author:
 * @Date: 2022-06-01 14:30:18
 * @LastEditTime: 2022-06-03 11:59:51
 */
package twentity

//leetcode 22.括号的生成   回朔算法 + 递归

//jd: 数字n代表生成括号的对数,用于能够生成所有可能的并且有效的括号组合.
// input: n=3  output: ["((()))","(()())","(())()","()(())","()()()"]
// input:
// input: n = 1 output: ["()"]

func generateParenthesis(n int) []string {
	box := new([]string)
	if n == 1 {
		return []string{"()"}
	}
	generateBackTracking("", n, n, box)
	return *box
}

func generateBackTracking(path string, left, right int, b *[]string) {

	if right == 0 {
		*b = append(*b, path)
		return
	}
	// 生成左括号
	if left > 0 {
		generateBackTracking(path+"(", left-1, right, b)
	}
	// 生成右括
	if right > left {
		generateBackTracking(path+")", left, right-1, b)
	}
}
