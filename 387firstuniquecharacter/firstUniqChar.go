package _87firstuniquecharacter


func firstUniqChar(s string) int {
		var lf [26]int
		for i, ch := range s {  //遍历重复值为一次性值
		lf[ch - 'a'] = i
	}
		for i, ch := range s {
		if i == lf[ch - 'a'] {  //第一次出现是否是最后一次索引
		return i //第一个直接返回
	} else {
		lf[ch - 'a'] = -1 // NOF FOUND
	}
	}
		return -1
	}



