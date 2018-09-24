package selectors

import (
	"errors"

	"../sorts"
)

func SelectFromSorted(numbers []int, k int) (int, error) {
	if k >= len(numbers) {
		return 0, errors.New("Índicie maior que o tamanho do vetor")
	} else if k <= 0 {
		return 0, errors.New("Índicie precisa ser maior que 0")
	}
	return numbers[k], nil
}

func SimpleSelect(numbers []int, k int) (int, error) {
	sorted := sorts.MergeSort(numbers)
	return SelectFromSorted(sorted, k)
}
