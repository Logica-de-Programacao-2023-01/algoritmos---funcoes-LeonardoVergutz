package main

import "fmt"

func segundo(inteiros []int) int {
	maior := inteiros[0]
	segundomaior := inteiros[1]

	if segundomaior > maior {
		maior, segundomaior = segundomaior, maior
	}

	for i := 2; i < len(inteiros); i++ {
		if inteiros[i] > maior {
			segundomaior = maior
			maior = inteiros[i]
		} else if inteiros[i] > segundomaior && inteiros[i] != maior {
			segundomaior = inteiros[i]
		}
	}
	return segundomaior
}
func main() {
	var numeros = []int{10, 230, 40, 42, 54, 234, 65, 34, 657}
	fmt.Print("O segundo maior valor do slice ", numeros, " é:", segundo(numeros))
}
