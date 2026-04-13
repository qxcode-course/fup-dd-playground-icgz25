package main

import "fmt"

func main() {
	var nome string
	var idade int

	fmt.Scanln(&nome)
	fmt.Scanln(&idade)

	var classificacao string

	if idade < 12 {
		classificacao = "crianca"
	} else if idade < 18 {
		classificacao = "jovem"
	} else if idade < 65 {
		classificacao = "adulto"
	} else if idade < 1000 {
		classificacao = "idoso"
	} else {
		classificacao = "mumia"
	}

	fmt.Printf("%s eh %s\n", nome, classificacao)
}