package main

import "fmt"

func verifica(inteiros []int, numero int) int {
	for i, num := range inteiros {
		if num == numero {
			return i
		}
	}
	return -1
}

func main() {
	var valor int
	var numeros = []int{10, 20, 34, 452, 53, 723, 20, 75}

	fmt.Print("Digite um número para verificar sua existencia no slice:")
	fmt.Scan(&valor)

	fmt.Println("A posição correspondente ao valor inserido é:", verifica(numeros, valor))
	fmt.Println(numeros)
}
