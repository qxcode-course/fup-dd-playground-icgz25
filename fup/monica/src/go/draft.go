package main
import "fmt"
func main() {
    var M, A, B int

    fmt.Scanln(&M)
    fmt.Scanln(&A)
    fmt.Scanln(&B)

    C := M - (A + B)

    maisVelho := A
    if B > maisVelho {
        maisVelho = B
    }
    if C > maisVelho {
        maisVelho = C
    }

    fmt.Println(maisVelho)
}
