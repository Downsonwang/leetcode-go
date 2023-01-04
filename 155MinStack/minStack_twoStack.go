package _55MinStack


type MingStack struct {
	dataStack []int
	minStack []int
}
func Constructor1() MingStack{
	return MingStack{}
}
func (this *MingStack) Push(val int){
	this.dataStack = append(this.dataStack,val)
	n := len(this.minStack)
	if n == 0 || this.minStack[n-1] > val {
		this.minStack = append(this.minStack,val)
	}else {
		this.minStack = append(this.minStack,this.minStack[n-1])

	}
}

func (this *MingStack) Pop(){
	this.dataStack = this.dataStack[:len(this.dataStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}
func (this *MingStack) Top() int {
	return this.dataStack[len(this.dataStack)-1]
}

func (this *MingStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}