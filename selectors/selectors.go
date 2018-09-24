package selectors

import (
	"errors"
	"log"

	"../sorts"
)

// O(n log n)

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

// LINEAR : O(n)

func LinearSelect(numbers []int, k int) int {
	return LinearSelectionRecursive(numbers, 0, len(numbers)-1, k)
}

func Partitions(numbers []int, l int, r int, pvt int) int {
	var tmp int
	pivot := numbers[pvt]
	numbers[pvt] = numbers[l]

	// int pivot = array[l];
	i := l + 1
	for (numbers[i] < pivot) && (i < r+1) {
		i++
	}

	jstart := i

	for j := jstart; j < r+1; j++ {
		if numbers[j] < pivot {
			tmp = numbers[i]
			numbers[i] = numbers[j]
			numbers[j] = tmp
			i++
		}
	}

	// swap pivot
	numbers[l] = numbers[i-1]
	numbers[i-1] = pivot
	return i - 1
}

func LinearSelectionRecursive(numbers []int, l int, r int, k int) int {
	var pivotPt int
	if l > r {
		log.Fatal("Invalid array index! Left index > Right index")
	} else if (k > r) || (k < l) {
		log.Fatal(k+1, "_th Statistic is outside array domain")
	}

	if l < r {
		pivotPt = (l + r) / 2
	} else {
		pivotPt = l
	}

	pst := Partitions(numbers, l, r, pivotPt)

	if pst == k {
		return numbers[pst]
	}

	if pst > k {
		return LinearSelectionRecursive(numbers, l, pst-1, k)
	}

	if pst < k {
		return LinearSelectionRecursive(numbers, pst+1, r, k)
	}

	return 0
}
