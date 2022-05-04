package main

type CQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() CQueue {
	return CQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 { //这里要注意,一定要为空的时候
		for len(this.inStack) != 0 {
			v := this.inStack[len(this.inStack)-1]
			this.outStack = append(this.outStack, v)
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
	}

	if len(this.outStack) == 0 {
		return -1
	}

	ret := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]

	return ret
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
