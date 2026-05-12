package main
import "fmt"
func main() {
    var A, B int
    fmt.Scanln(&A)
    fmt.Scanln(&B)

    fmt.Print("[ ")

    for i, j := A, B; i < j; i, j = i+1, j-1 {
        fmt.Printf("%d %d ", i, j)
    }

    for i, j := (B-1), (A+1); i >= A; i, j = i-1, j+1 {
        fmt.Printf("%d %d ", i, j)
    }

    fmt.Println("]")
}