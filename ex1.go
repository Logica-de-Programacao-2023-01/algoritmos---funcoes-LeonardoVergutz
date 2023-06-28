package main

import "fmt"

func divisao(slice [5]int) int {
	var soma, quantidade, media int

	for _, num := range slice {
		soma += num
		quantidade++
	}
	media = soma / quantidade
	return media
}

func main() {
	var inteiros = [5]int{10, 20, 30, 40, 50}
	fmt.Print("A soma do slice", inteiros, " é:", divisao(inteiros))
}
