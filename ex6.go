package main

import (
	"errors"
	"fmt"
)

func concate(str []string) (string, error) {
	var soma string
	if len(str) == 0 {
		return "", errors.New("Slice vazio!")
	}
	soma = str[0]
	for _, num := range str[1:] {
		soma += "," + num
	}
	return soma, nil
}

func main() {
	var strs = []string{"malandragem", "Sacanagem", "maluquice", "Doidera", "barbaridade"}

	concatenação, err := concate(strs)
	if err != nil {
		fmt.Println("Erro", err)
		return
	}
	fmt.Print("A concatenação do slice ", strs, " é: ", concatenação)
}
