package _87firstuniquecharacter


// 使用位掩码记录 为了超过 100%
// 一个记录是否某个字符是否已经出现过 或运算 1 ｜ 0 = 0  1 ｜1 =1
// 另外一个变量记录某个字符出现过一次以上

func firstUniqueChar(s string) int{
	var exist,repeated int32
	for _, v := range s {
		tmp := int32(1 << (v - 'a'))
		if exist&tmp > 0 {
			repeated |= tmp
		}else {
			exist |= tmp
		}
	}

	for i, v := range s {
		if repeated&(1<<(v-'a')) == 0 {
			return i
		}
	}
	return -1
}