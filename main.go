package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"./selectors"
	"./sorts"
)

func main() {

	// Verifica quantos argumentos entraram:
	if len(os.Args) < 3 {
		fmt.Println("Faltando parametros de entrada!\n\n")
		fmt.Println("1 - Tamanho do vetor de entrada\n")
		fmt.Println("2 - Indicie k-menor\n")
		return
	}

	nString := os.Args[1]
	kString := os.Args[2]
	var err error
	var n int64
	var k int64

	n, err = strconv.ParseInt(nString, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	k, err = strconv.ParseInt(kString, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	numbers := generateSlice(int(n))
	fmt.Println("\nArray desordenado:\n\n", numbers)
	fmt.Println("\nArray ordenado:\n\n", sorts.MergeSort(numbers), "\n")

	kSelection, err := selectors.SimpleSelect(numbers, int(k))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("\nNumero encontrado:\n\n", kSelection)

}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		// slice[i] = rand.Intn(99999) - rand.Intn(99999)
		slice[i] = rand.Intn(10) //- rand.Intn(10)
	}
	return slice
}
