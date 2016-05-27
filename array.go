package capable

const capacity = 1024

func array() [capacity]int {
	var a [capacity]int
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}

	return a
}

func sliceLimit() []int {
	a := make([]int, capacity)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	return a
}
func slice() []int {
	a := make([]int, 0)
	for i := 0; i < capacity; i++ {
		a = append(a, 1)
	}
	return a
}
