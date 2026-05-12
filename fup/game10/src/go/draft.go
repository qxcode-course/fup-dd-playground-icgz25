package main
import ("fmt")
func main() {
    var N, D, A int

    fmt.Scanln(&N)
    fmt.Scanln(&D)
    fmt.Scanln(&A)

    apertos := (D - A + N) % N

    fmt.Println(apertos)
}
