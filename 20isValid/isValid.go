package _0isValid

// 20. 有效的括号
// 给定一个包括 '(',')'，'{','}','[',']'的字符串s,判断字符串是否有效
// 有效字符串需满足:
// 1.左括号必须用相同类型的右括号闭合
// 2.左括号必须以正确的顺序闭合
// 3.每个右括号都有一个对应相同类型的左括号

func isValid(s string) bool {
	// 奇数为false
	if len(s) == 1 {
		return false
	}
	var stack []string

	// 存储特征
	storageFeature := map[string]string{}
	storageFeature["("] = ")"
	storageFeature["{"] = "}"
	storageFeature["["] = "]"
	for _,v := range s {
		c := string(v)
		if _,ok := storageFeature[c]; ok {
			stack = append(stack,c) //压栈
		} else {
			//空栈
			if len(stack) == 0 {
				return false
			}
			// 获取栈顶元素
			top := stack[len(stack) - 1]

			// 括号匹配,才能出栈
			if v,ok := storageFeature[top]; ok && v==c {
				//弹栈
				stack = stack[0 : len(stack) - 1]
			}else {
				return false
			}
		}
	}

   return  len(stack) == 0
}

