package main
import "fmt"
func main() {
    var A, B int

    fmt.Scanln(&A)
    fmt.Scanln(&B)

    for i := A; i < B; i++ {
        fmt.Println(i)
    }
}