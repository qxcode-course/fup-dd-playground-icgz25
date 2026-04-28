package main
import "fmt"
func main() {
	var p, d1, d2 int

	fmt.Scan(&p, &d1, &d2)

	soma := d1 + d2

	var vencedor int

	if soma%2 == 0 {
		if p == 0 {
			vencedor = 0
		} else {
			vencedor = 1
		}
	} else {
		if p == 0 {
			vencedor = 1
		} else {
			vencedor = 0
		}
	}
	fmt.Println(vencedor)
}