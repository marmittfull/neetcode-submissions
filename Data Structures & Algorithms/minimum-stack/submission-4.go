type MinStack struct {
	elements []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		elements: []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
	this.elements = append(this.elements, val)
}

func (this *MinStack) Pop() {
	val := this.elements[len(this.elements) - 1]
	this.elements = this.elements[:len(this.elements) - 1]
	if val == this.minStack[len(this.minStack) - 1] {
		this.minStack = this.minStack[:len(this.minStack) - 1]
	}
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements) - 1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack) - 1]
}
