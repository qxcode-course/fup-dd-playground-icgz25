package main
import "fmt"
func main() {
    var A, B int

    fmt.Scanln(&A, &B)

    fmt.Print("[ ")
    for i := A; i < B; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println("]")
}
