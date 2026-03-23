package main
import "fmt"
func main() {
    var c float64

    fmt.Scan(&c)

    f := c*9/5 + 32

    fmt.Printf("%.6f\n", f)
}
