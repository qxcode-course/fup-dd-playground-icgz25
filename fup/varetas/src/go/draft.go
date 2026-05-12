package main
import "fmt"
func main() {
    var a, b, c int

    fmt.Scanln(&a)
    fmt.Scanln(&b)
    fmt.Scanln(&c)

    if a < b+c && b < a+c && c < a+b {
        fmt.Println("True")
    } else {
        fmt.Println("False")
    }
}
