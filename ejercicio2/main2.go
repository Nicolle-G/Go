package main

import "fmt"

const size = 5

type RingBuffer struct {
	tail   int
	head   int
	count  int
	buffer [size]int
}

func (rb *RingBuffer) append(value int) {
	if rb.count == size {
		fmt.Println("Buffer full. Overwriting the oldest value.")
		rb.head = (rb.head + 1) % size
	} else {
		rb.count++
	}

	rb.buffer[rb.tail] = value
	rb.tail = (rb.tail + 1) % size
}

func (rb *RingBuffer) pop() int {
	if rb.count == 0 {
		fmt.Println("Buffer empty. Nothing to read.")
		return -1
	}

	valor := rb.buffer[rb.head]
	rb.head = (rb.head + 1) % size
	rb.count--
	return valor
}

func (rb *RingBuffer) PrintBuffer() {
	fmt.Println("buffer content")
	for i := 0; i < size; i++ {
		fmt.Printf("[%d] ", rb.buffer[i])
	}
	fmt.Println()
	fmt.Println("Head:", rb.head, "Tail:", rb.tail, "Count:", rb.count)
}

func (rb *RingBuffer) values() []int {
	vals := make([]int, 0, rb.count)
	for i, j := 0, rb.head; i < rb.count; i++ {
		vals = append(vals, rb.buffer[j])
		j = (j + 1) % size
	}
	return vals
}

func main() {

	var rb RingBuffer

	for i := 1; i <= 7; i++ {
		fmt.Println("inserting", i)
		rb.append(i)
		rb.PrintBuffer()
	}

	fmt.Println("Taking out 2 values with pop()")
	fmt.Println("Extracted value:", rb.pop())
	rb.PrintBuffer()
	fmt.Println("Extracted value:", rb.pop())
	rb.PrintBuffer()

}
