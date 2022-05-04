package main

type MyQueue struct {
	instack  []int
	outstack []int
}

func Constructor() MyQueue {
	return MyQueue{
		instack:  make([]int, 0),
		outstack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.instack = append(this.instack, x)
}

func (this *MyQueue) in2out() {
	for len(this.instack) > 0 {
		num := this.instack[len(this.instack)-1]
		this.instack = this.instack[:len(this.instack)-1]
		this.outstack = append(this.outstack, num)
	}
}

func (this *MyQueue) Pop() int {
	if len(this.outstack) == 0 { //这里要注意,一定要为空的时候
		this.in2out()
	}

	num := this.outstack[len(this.outstack)-1]
	this.outstack = this.outstack[:len(this.outstack)-1]

	return num
}

func (this *MyQueue) Peek() int {
	if len(this.outstack) == 0 {
		this.in2out()
	}

	return this.outstack[len(this.outstack)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.outstack) == 0 {
		this.in2out()
	}

	if len(this.outstack) == 0 {
		return true
	}

	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
