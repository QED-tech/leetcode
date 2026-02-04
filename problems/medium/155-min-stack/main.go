package main

import "fmt"

type stackElem struct {
	val int
	min int
}

type MinStack struct {
	elems []stackElem
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.elems) == 0 {
		this.elems = append(this.elems, stackElem{val: val, min: val})
		return
	}

	el := stackElem{
		val: val,
		min: min(this.GetMin(), val),
	}
	this.elems = append(this.elems, el)
}

func (this *MinStack) Pop() {
	this.elems = this.elems[:len(this.elems)-1]
}

func (this *MinStack) Top() int {
	return this.elems[len(this.elems)-1].val
}

func (this *MinStack) GetMin() int {
	return this.elems[len(this.elems)-1].min
}

func main() {
	ms := Constructor() // MinStack

	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)

	fmt.Println(ms.GetMin()) // -> -3

	ms.Pop()

	fmt.Println(ms.Top())    // -> 0
	fmt.Println(ms.GetMin()) // -> -2
}
