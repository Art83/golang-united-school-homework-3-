package homework

import "fmt"

func average(input [15]float32) (result float32) {
	var avg float32
	var cnt int
	fmt.Println(len(input))
	for _, v := range input {
		avg += v
		cnt++
	}
	return avg / float32(cnt)
}
