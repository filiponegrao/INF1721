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
		fmt.Println("ERRO: Faltando parametros de entrada!\n\n")
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

	if k > n {
		fmt.Println("ERRO: k nao pode ser maior que n\n")
		return
	}
	if k <= 0 {
		fmt.Println("ERRO: k nao pode ser menor ou igual a 0\n")
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
	fmt.Println("\nNumero encontrado MergeSelect:\n\n", kSelection)

	kSelectionLinear := selectors.LinearSelect(numbers, int(k))

	fmt.Println("\nNumero encontrado Linear Select:\n\n", kSelectionLinear)

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
