package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	vals := make([]string, len(input))
	for i, _ := range input {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	for i, k := range keys {
		vals[i] = input[k]
	}

	return vals
}
