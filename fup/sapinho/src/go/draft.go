package main
import "fmt"
func main() {
	var p, s, d int

	fmt.Scan(&p, &s, &d)

	for {
		p -= s

		if p <= 0 {
			fmt.Println(p+s, "saiu")
			break
		}

		fmt.Println(p, p+s)

		p += d
	}
}