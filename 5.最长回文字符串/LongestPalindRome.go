package __最长回文字符串

/*
   输入 s = "babad"
   输出： "bab"
   输入 s = "cbbd"
   输出 "bb"
    两边走


  中心扩散，  先找出和中心不想等的左右索引 这样可以忽略奇偶长度对称性的差异
  下一次搜索直接将索引移动到上一次的初始右索引，节省重复计算
   Todo:其实全部循环完成后再生成子字符串就行了
*/
func LongestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	var maxLen = 0
	var res = ""
	index := 0
	for index < len(s) {
		var left = index
		var right = index
		for left >= 0 && s[left] == s[index] {
			left--
		}
		for right < len(s) && s[right] == s[index] {
			right++
		}
		var nextPoint = right
		for left >= 0 && right < len(s) && s[right] == s[left] {
			left--
			right++
		}
		var tmpMaxLen = right - left - 1
		if tmpMaxLen > maxLen {
			res = s[left+1 : right]
			maxLen = tmpMaxLen
		}
		index = nextPoint
	}
	return res
}
