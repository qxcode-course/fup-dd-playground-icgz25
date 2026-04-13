package main

import "fmt"

func main() {
	var m, a, b int

	fmt.Scan(&m, &a, &b)

	c := m - a - b

	maisVelho := a

	if b > maisVelho {
		maisVelho = b
	}
	
	if c > maisVelho {
		maisVelho = c
	}

	fmt.Println(maisVelho)
}