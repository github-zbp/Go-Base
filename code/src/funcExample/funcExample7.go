package funcExample

func Add(vals ...int) int {
	var sum int
	for _, val := range vals {
		sum += val
	}

	return sum
}