package main

import (
	"errors"
	"fmt"
	"sort"
)

func crescente(nums []int) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("O slice esta vazio")
	}

	sortedSlice := make([]int, len(nums))
	copy(sortedSlice, nums)
	sort.Ints(sortedSlice)
	return sortedSlice, nil
}

func main() {
	var numeros = []int{10, 230, 34, 45, 56, 3, 2, 345, 34, 68, 987, 232, 19, 8, 1}

	crescente, err := crescente(numeros)

	if err != nil {
		fmt.Print("Erro:", err)
	}
	fmt.Println(numeros)
	fmt.Println(crescente)
}
