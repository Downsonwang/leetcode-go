package __Z字数变换

import "fmt"

func convert(s string, numRows int) string {
	var res string
	data := make([]string, numRows)
	// 初始位置
	curPos := 0
	// 标记是否转向
	shouldTurn := -1
	if "" == s || numRows < 1 {
		return res
	}
	if numRows == 1 {
		return s
	}
	for _, value := range s {
		// 添加进tmp 位置为curPos
		data[curPos] += string(value)
		fmt.Println(data)
		if curPos == 0 || curPos == numRows-1 {
			shouldTurn = -shouldTurn
		}
		curPos += shouldTurn
	}
	for _, v := range data {
		res += v
	}

	// fmt.Println(data)
	return res
}
