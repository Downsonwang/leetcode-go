/*
 * @Descripttion:  最长公共前缀
 * @Author: downsonwang
 * @Date: 2022-09-07 09:16:48
 * @LastEditTime: 2022-09-07 14:28:50
 */
package longestcommonprefix

// leetcode 最长公共前缀
func Longestcommonprefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]

	for _, v := range strs {
		prefix = strength(prefix, v)

		if len(prefix) == 0 {
			break
		}

	}
	return prefix
}

//拿出最短的公共前缀
func strength(str1, str2 string) string {
	l := min(len(str1), len(str2))
	index := 0
	for index < l && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
