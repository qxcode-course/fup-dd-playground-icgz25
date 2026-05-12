package main
import ("fmt")
func main() {
    var C, M int
    fmt.Scanln(&C)

    passageiros := 0

    for {
        fmt.Scanln(&M)
        passageiros += M

        if passageiros == 0 {
            fmt.Println("vazio")
        } else if passageiros < C {
            fmt.Println("ainda cabe")
        } else if passageiros == C {
            fmt.Println("lotado")
        } else if passageiros >= 2*C {
            fmt.Println("hora de partir")
            break
        } else {
            fmt.Println("lotado")
        }
    }
}