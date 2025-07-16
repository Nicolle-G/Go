package main

import (
	"testing"
)

func TestRingBufferInsertAndPop(t *testing.T) {

	var rb RingBuffer

	for i := 1; i <= 7; i++ {
		rb.append(i)
	}

	got := rb.values()
	want := []int{3, 4, 5, 6, 7}

	if !equal(got, want) {
		t.Errorf("Expected %v but it was obtained %v", want, got)
	}

	val1 := rb.pop()
	val2 := rb.pop()

	if val1 != 3 || val2 != 4 {
		t.Errorf("Error in pop expected 3 and 4 and got %d y %d", val1, val2)
	}

}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
