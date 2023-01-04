package morewater

// 找到最小的
func minNum(fir, sec int) int {
	if fir > sec {
		return sec
	}
	return fir
}

// 找到最大的
func maxNum(fir, sec int) int {
	if fir > sec {
		return fir
	}
	return sec
}

// 核心算法  双指针
func maxArea(height []int) int {
	i := 0               // 右移游标
	j := len(height) - 1 //左移游标
	area := 0            // 最终结果
	final := 0
	if len(height) < 2 {
		return 0
	}
	// 面积 = 以最短的高度 * 两板的距离
	for i < j {

		if height[i] < height[j] {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}
		if final < area {
			final = area
		}

	}
	return final
}
