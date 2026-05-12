package main
import "fmt"
func main() {
    var N int
    var inicio string
    fmt.Scanln(&N)
    fmt.Scanln(&inicio)

    fmt.Print("[ ")

    pe := inicio
    for i := 0; i <= 10; i++ {
        if i == N {
            continue
        }

        if i == 10 {
            fmt.Print("ceu ")
            break
        }

        fmt.Printf("%d%s ", i, pe)

        if pe == "d" {
            pe = "e"
        } else {
            pe = "d"
        }
    }

    fmt.Println("]")
}