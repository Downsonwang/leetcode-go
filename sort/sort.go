package sort


/*
  冒泡排序 插入排序 选择排序
 */
// 冒泡排序 只会操作相邻的两个数据, 每次冒泡操作都会对相邻的两个元素进行比较,看是否满足大小关系要求 . 如果不满足就让它俩互换.  一次冒泡会让至少一个元素移动到它应该在的位置
// 重复n次 就完成了n个数据的排序工作
func BubbleSort(a []int,n int ){
   if n <= 1 {
	   return
   }
   for i := 0 ; i < n ; i++ {
	   flag := false //提前退出标志
	   for j := 0; j < n-i-1;j++{
		   if a[j] > a[j+1]{
			   a[j],a[j+1] = a[j+1],a[j]
			   flag = true
		   }
	   }
	   //如果没有交换数据 提前退出
	   if !flag{
		   break
	   }
   }
}

// 插入排序 a表示数组 n表示数组大小
// 只要遍历数组,找到数据应该插入的位置将其插入即可
// 已排序期间和未排序期间 数组第一个元素不用遍历
// 核心思想是 取未排序的元素 在已排序区间中找到合适的插入位置将其插入
// 元素比较 + 元素移动
func InsertionSort(a []int,n int){
	if n <= 1 {
		return
	}
	for i :=1; i < n ; i++{
		val := a[i]
		j := i - 1
		// 寻找要插入的位置并移动数据
		for ; j >= 0 ; j-- {
			if a[j] > val{
				a[j+1] = a[j]
			}else {
				break
			}
		}
		a[j+1] = val
	}
}

// 选择排序
// 选择排序每次会从未排序区间中找到最小的元素,将其放到已排序区间的末尾
func SelectionSort(a []int, n int ){
	if n <= 1 {
		return
	}
	for i := 0; i < n ; i++ {
		//查找最小值
		min := i
		for j := i + 1; j < n;j++{
			if a[j] < a[min]{
				min = j
			}
		}
		// 交换
		a[i],a[min] = a[min],a[i]
	}
}