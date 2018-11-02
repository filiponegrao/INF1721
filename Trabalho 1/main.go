package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/filiponegrao/INF1721/selectors"
	"github.com/filiponegrao/INF1721/sorts"
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
		fmt.Println("ERRO: k nao pode ser maior que n")
		return
	}
	if k <= 0 {
		fmt.Println("ERRO: k nao pode ser menor ou igual a 0")
		return
	}

	numbers := generateSlice(int(n))
	prinStart()
	fmt.Println("\nArray desordenado:\n\n", numbers)
	fmt.Println("\nArray ordenado:\n\n", sorts.MergeSort(numbers))

	fmt.Println("\n#### SORT SELECTION:")
	start := time.Now()
	sortSelection, _ := selectors.SimpleSelect(numbers, int(k))
	elapsed := time.Since(start)
	fmt.Println("* Tempo de execucao: ", elapsed)
	fmt.Println("* Valor retornado: ", sortSelection)

	fmt.Println("\n#### LINEAR SELECTION:")
	start = time.Now()
	linearSelection := selectors.LinearSelect(numbers, int(k))
	elapsed = time.Since(start)
	fmt.Println("* Tempo de execucao: ", elapsed)
	fmt.Println("* Valor retornado: ", linearSelection)

}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		// slice[i] = rand.Intn(99999) - rand.Intn(99999)
		slice[i] = rand.Intn(10000) //- rand.Intn(10)
	}
	return slice
}

func prinStart() {
	fmt.Println("##########################################")
	fmt.Println("#### [INF1721] Analise de Algoritimos ####")
	fmt.Println("##########################################")
	fmt.Println("\n# Comparacao de selecao com SortSelection e")
	fmt.Println("# LinearSelection")

}
