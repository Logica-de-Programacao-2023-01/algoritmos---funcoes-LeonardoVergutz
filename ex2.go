package main

import (
	"fmt"
	"strings"
)

func vogais(str string) int {
	count := 0
	letras := "A ,E ,I ,O ,U a, e, i, o, u"

	for _, char := range str {
		if strings.ContainsRune(letras, char) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Print("A quantidade de vogais presentes na string (siyfowhafhfuh) é:", vogais("siyfowhafhfuh"))
}
