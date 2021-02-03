package sort

import (
	"fmt"
	"log"
	"time"
)

type BinaryHeap struct {
	size int
	arr  []int
}

func (bh *BinaryHeap) Swim(n int) {
	for n/2 >= 1 && bh.arr[n] > bh.arr[n/2] {
		// 大于父节点
		bh.arr[n], bh.arr[n/2] = bh.arr[n/2], bh.arr[n]
		n = n / 2
	}
}

func (bh *BinaryHeap) Sink(n int) {
	parent := n
	for parent*2 < bh.size {
		firstChild := parent * 2
		maxChild := firstChild
		if firstChild < bh.size && bh.arr[firstChild] < bh.arr[firstChild+1] {
			maxChild = firstChild + 1
		}
		if bh.arr[parent] > bh.arr[maxChild] {
			break
		}
		bh.arr[parent], bh.arr[maxChild] = bh.arr[maxChild], bh.arr[parent]
		parent = maxChild
	}
}

func (bh *BinaryHeap) DeleteMax() int {
	max := bh.arr[1]
	bh.arr[1], bh.arr[bh.size] = bh.arr[bh.size], bh.arr[1]
	bh.size -= 1
	bh.Sink(1)
	return max
}

func (bh *BinaryHeap) Insert(value int) {
	bh.size += 1
	if len(bh.arr) > bh.size {
		// 容量足够
		print("容量足够")
		bh.arr[bh.size] = value
	} else {
		// 容量不够了
		bh.arr = append(bh.arr, value)
	}
	bh.Swim(bh.size)
}

func (bh *BinaryHeap) GetArr() []int {
	return bh.arr
}

func (bh *BinaryHeap) GetSize() int {
	return bh.size
}

func NewBinaryHeap(arr []int) BinaryHeap {
	arr = append([]int{0}, arr...)
	size := len(arr) - 1
	return BinaryHeap{size: size, arr: arr}
}

func (bh *BinaryHeap) Sort() {
	// 下沉法让堆有序
	start := time.Now()
	size := bh.GetSize()
	for n := (size + 1) / 2; n >= 1; n-- {
		bh.Sink(n)
	}
	// 删除法排序
	for bh.GetSize() > 0 {
		bh.DeleteMax() // 并没有删除，而是放在最后一个元素
	}
	spend := time.Since(start)
	log.Println(fmt.Sprintf("堆排序花费%v", spend))

}
