package homework

func reverse(input []int64) (result []int64) {
	answ := make([]int64, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		answ[len(input)-i-1] = input[i]
	}
	return answ
}
