package main

import "fmt"

func conc(str [5]string) string {
	var soma string
	for _, num := range str {
		soma += num
	}
	return soma
}

func main() {
	var strs = [5]string{"malandragem", "Sacanagem", "maluquice", "Doidera", "barbaridade"}
	fmt.Print("A concatenação do slice ", strs, "é: ", conc(strs))

}
