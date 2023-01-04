package __无重复字符的最长子串

import "fmt"

// 滑动弹窗
func LengthOfLongestSubstring(s string) int {
	// is := make([]map[int]int32,0)
	/*
		var is [128]int
		 if len(s) <= 1 {
			return len(s)
		}
	   max := 1

	    is[0] = 1 //边界
	    for i := 1 ; i < len(s) ; i++ {
	    	for j := i - is[i - 1] ; j < i ; j++{
	             if is[i] == is[j]{
	             	 is[i] = i - j
	             	 goto LOOP
				 }
			}
			is[i]=is[i-1]+1   //否则就长度+1
		LOOP:
			if is[i]>max {
				max = is[i]
			}

		}
		fmt.Println(max)
		return  max
		*
	*/
	if len(s) <= 1 {
		return len(s)
	}
	var is [256]int
	left := 0
	right := 0
	wins := 0

	for i := 0; i < len(is); i++ {
		is[i] = -1
	}
	for right < len(s) {
		fmt.Println(is[s[right]-'a'])
		if idx := is[s[right]-'a']; idx != -1 && idx >= left {
			left = idx + 1
		}
		tmp := right - left + 1
		if tmp > wins {
			wins = tmp
		}
		is[s[right]-'a'] = right
		right++
	}
	return wins
}
