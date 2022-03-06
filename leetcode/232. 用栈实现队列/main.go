package main

type MyQueue struct {
	list []int
}

func Constructor() MyQueue {
	return MyQueue{
		list: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.list = append(this.list, x)
}

func (this *MyQueue) Pop() int {
	data := this.list[0]
	this.list = this.list[1:]
	return data
}

func (this *MyQueue) Peek() int {
	return this.list[0]
}

func (this *MyQueue) Empty() bool {
	if len(this.list) == 0 {
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
