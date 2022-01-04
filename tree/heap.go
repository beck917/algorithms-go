package main

type Heap struct {
	arr   []int
	n     int
	count int //堆目前的长度
}

func New(capacity int) *Heap {
	return &Heap{
		arr:   make([]int, capacity),
		count: 0,
		n:     10,
	}
}

// 核心逻辑
// 数组从1开始，方便用数组存储tree 节点，
// 这样父节点就是i/2，比如新插入的几点为2或者3，那么他的父节点就是1
// 子节点就是i*2和i*2+1
// 插入逻辑是从最后一个子节点插入，然后向上冒泡，对比和父节点的大小，如果到达root节点或者比父节点小，则停止遍历
func (heap *Heap) insert(data int) {
	if heap.count >= heap.n {
		return
	}
	heap.count++
	heap.arr[heap.count] = data
	i := heap.count
	for i/2 > 0 && heap.arr[i/2] < heap.arr[i] {
		heap.arr[i/2], heap.arr[i] = heap.arr[i], heap.arr[i/2]
		i = i / 2
	}
	return
}

// 删除堆顶算法，也就是常用的取top 10的算法
// 删除堆顶元素，实际上是将堆顶元素和堆底元素进行交换
// 移除堆顶元素，然后将堆底元素放到堆顶，然后向下冒泡
func (heap *Heap) removeMax() {
	heap.count--
	//heap.arr = heap.arr[1:]
	heap.arr[1] = heap.arr[heap.count]
	heap.arr[heap.count] = 0
	heap.heapify(1)
}

func (heap *Heap) heapify(i int) {
	for {
		maxpos := i
		if i*2 < heap.count && heap.arr[i*2] > heap.arr[i] {
			heap.arr[i*2], heap.arr[i] = heap.arr[i], heap.arr[i*2]
			maxpos = i * 2
		}

		if i*2+1 < heap.count && heap.arr[i*2+1] > heap.arr[i] {
			heap.arr[i*2+1], heap.arr[i] = heap.arr[i], heap.arr[i*2+1]
			maxpos = i*2 + 1
		}

		if maxpos == i {
			break
		}

		i = maxpos
	}
}
