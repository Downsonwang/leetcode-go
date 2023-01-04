package romenumtoint

func romanToInt(s string) int {
	fin := 0
	option := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	length := len(s)
	for i := 0; i < length; i++ {
		// 当前小于右边 就相当于"I"遇到了"V"
		if i < length-1 && option[s[i]] < option[s[i+1]] {
			// 给 -
			fin -= option[s[i]]
		} else {
			// 大于 直接给答案
			fin += option[s[i]]
		}
	}
	return fin
}
