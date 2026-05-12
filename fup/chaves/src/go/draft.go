package main
import "fmt"
func main() {
    var n int

    fmt.Scanln(&n)

    if n > 0 {
        fmt.Println("positivo")
    } else if n == 0 {
        fmt.Println("nulo")
    } else {
        fmt.Println("negativo")
    }
}
