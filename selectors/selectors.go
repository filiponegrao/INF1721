package selectors

import (
	"errors"

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
	// return LinearSelectionRecursive(numbers, 0, len(numbers)-1, k)
	return Select(numbers, k)
}

func Select(numbers []int, k int) int {
	n := len(numbers)
	if n == 1 {
		return numbers[0]
	}
	parts := SplitNumbersInParts(numbers, 5) // O(n)
	for i := 0; i < len(parts); i++ {        // O(n/5)
		parts[i] = SmallArraySort(parts[i]) // O(5)
	}
	var medians []int
	for i := 0; i < len(parts); i++ { // O(n)
		partMedian := GetMedian(parts[i]) // cst
		medians = append(medians, partMedian)
	}
	median := Select(medians, len(medians)/2)

	left, equals, right := CreateArrayOfLessEqualsAndGrathers(numbers, median) // O(n)

	if k < len(left) {
		return Select(left, k)
	} else if k <= (len(left) + len(equals)) {
		return median
	} else {
		return Select(right, k-len(left)-len(equals))
	}
}

func CreateArrayOfLessEqualsAndGrathers(numbers []int, value int) ([]int, []int, []int) {
	var less []int
	var equals []int
	var grater []int

	for i := 0; i < len(numbers); i++ {
		if numbers[i] < value {
			less = append(less, numbers[i])
		} else if numbers[i] == value {
			equals = append(equals, numbers[i])
		} else {
			grater = append(grater, numbers[i])
		}
	}
	return less, equals, grater
}

func CreateArrayWithValuesLessThen(numbers []int, value int) []int {
	var results []int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < value {
			results = append(results, numbers[i])
		}
	}
	return results
}

func CreateArrayWithValuesMajorThen(numbers []int, value int) []int {
	var results []int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > value {
			results = append(results, numbers[i])
		}
	}
	return results
}

func CreateArrayWithValuesEqualTo(numbers []int, value int) []int {
	var results []int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == value {
			results = append(results, numbers[i])
		}
	}
	return results
}

func GetMedian(numbers []int) int {
	n := len(numbers)
	if n == 0 {
		return 0
	} else if n == 1 {
		return numbers[0]
	} else {
		return numbers[n/2]
	}
}

func SmallArraySort(numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			temp := numbers[i+1]
			numbers[i+1] = numbers[i]
			numbers[i] = temp
		}
	}
	return numbers
}

func SplitNumbersInParts(numbers []int, partAmount int) [][]int {
	var results [][]int
	n := len(numbers)
	parts := n / partAmount
	if n%partAmount > 0 {
		parts += 1
	}
	for i := 0; i < parts-1; i++ {
		slice := numbers[i*partAmount : (i*partAmount)+partAmount]
		results = append(results, slice)
	}
	var slice []int
	for i := ((parts - 1) * partAmount); i < n; i++ {
		slice = append(slice, numbers[i])
	}
	results = append(results, slice)
	return results
}
