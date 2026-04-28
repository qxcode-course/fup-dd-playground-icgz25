package main
import "fmt"
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	var meio int

	if (a > b && a < c) || (a < b && a > c) {
		meio = a
	} else if (b > a && b < c) || (b < a && b > c) {
		meio = b
	} else {
		meio = c
	}
    
	fmt.Println(meio)
}