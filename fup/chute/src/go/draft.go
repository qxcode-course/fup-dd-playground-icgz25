package main
import "fmt"
func main() {
    var produto, chute1, chute2 int

    fmt.Scanln(&produto)
    fmt.Scanln(&chute1)
    fmt.Scanln(&chute2)

    dist1 := abs(produto - chute1)
    dist2 := abs(produto - chute2)

    if dist1 < dist2 {
        fmt.Println("primeiro")
    } else if dist2 < dist1 {
        fmt.Println("segundo")
    } else {
        fmt.Println("empate")
    }
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
