package _87firstuniquecharacter


// 根据哈希标解决问题. 字符串中的第一个唯一字符
// 哈希表法 键为字符串的字符,值有两种情况:
// 若值为-1,则代表该元素在字符串中存在重复值
// 若值不为-1 则代表该元素在字符串的下标
// map不是顺序存储,而是随机存储,所以不能遍历map数据结构 然后返回第一个value!=-1的下标,因为它很可能不是第一个不重复的字符

func firstUniqCharr(s string) int {
   hashMap := make(map[string]int)
   for i,ch := range s {
	   if _,ok := hashMap[string(ch)];ok{
		   hashMap[string(ch)] = -1 //标志该字符元素弃用
	   }else{
		   hashMap[string(ch)] = i //该字符元素需使用,值直接记录下其对应的下标
	   }
   }
   for i,ch :=range s{
	   if hashMap[string(ch)] != -1 {
		   return i
	   }
   }

  return -1
}