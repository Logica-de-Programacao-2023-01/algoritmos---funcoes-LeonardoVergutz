package main

import (
	"errors"
	"fmt"
)

func pares(ints []int) ([]int, error) {
	var parez []int

	if len(ints) == 0 {
		return nil, errors.New("O slice esta vazio!")
	}
	for _, num := range ints {
		if num%2 == 0 {
			parez = append(parez, num)
		}
	}
	return parez, nil
}

func main() {
	var numeros = []int{10, 23, 345, 456, 3, 4, 764, 344, 734, 34, 82, 4677, 345, 74}

	pares, err := pares(numeros)

	if err != nil {
		fmt.Print("Erro:", err)
		return
	}
	fmt.Println(numeros)
	fmt.Println("Os numeros pares do slice acima são:", pares)
}
