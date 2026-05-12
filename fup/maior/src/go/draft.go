package main
import "fmt"
func main() {
    var chute1, valorReal float64
    var escolha string

    fmt.Scanln(&chute1)
    fmt.Scanln(&escolha)
    fmt.Scanln(&valorReal)

    if chute1 == valorReal {
        fmt.Println("primeiro")
    } else if escolha == "M" && valorReal > chute1 {
        fmt.Println("segundo")
    } else if escolha == "m" && valorReal < chute1 {
        fmt.Println("segundo")
    } else {
        fmt.Println("primeiro")
    }
}
