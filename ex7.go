package main

import "fmt"

func mt3(numeros int) int {
	numeros *= 3
	return numeros
}

func att(nums []int, multiplicador int) []int {
	resultado := make([]int, len(nums))

	for i, num := range nums {
		resultado[i] = mt3(num)
	}
	return resultado
}

func main() {
	var valores = []int{10, 20, 340, 543, 23, 65, 34, 76, 32}
	fmt.Println(valores)
	multiplicador := 3
	fmt.Println(att(valores, multiplicador))
}
