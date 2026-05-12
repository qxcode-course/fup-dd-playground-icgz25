package main
import "fmt"
func main() {
    var A, B int
    fmt.Scanln(&A, &B)

    fmt.Print("[ ")

    i := A
    for {
        if i >= B {
            break
        }

        if i%2 == 0 {
            i++
            continue
        }

        fmt.Print(i, " ")
        i++
    }

    fmt.Println("]")
}