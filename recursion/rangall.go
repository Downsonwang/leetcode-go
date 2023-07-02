package recusion


//实现一组数据集合的全排列
type RangeType struct {
	value []int
}

func NewRangeArray(n int) *RangeType{
	return &RangeType{
		value: make([]int,n),
	}
}

func (this *RangeType) RangeAll(start int){
	len := len(this.value)
	for i:=start; i<len;i++{
		// i = start 时输出自己
		// 如果i和start的值相同就没有必要交换
		if i == start || this.value[i] != this.value[start]{
			//交换当前这个与后面的位置
			this.value[i],this.value[start] = this.value[start],this.value[i]
			//递归处理索引+1
			this.RangeAll(start+1)
			//换回来 因为是递归 如果不换回来会影响后面的操作 并且出现重复
			this.value[i],this.value[start] = this.value[start],this.value[i]
		}
	}
}


