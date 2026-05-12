package main
import "fmt"
func main() {
    var L, R, D int

    fmt.Scanln(&L)
    fmt.Scanln(&R)
    fmt.Scanln(&D)
    if R > 50 && L < R && R > D {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}
