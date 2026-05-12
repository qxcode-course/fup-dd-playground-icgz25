package main
import "fmt"
func main() {
    var H, P, F, D int
    fmt.Scanln(&H, &P, &F, &D)

    pos := F
    for {
        if pos == H {
            fmt.Println("S")
            break
        }
        if pos == P {
            fmt.Println("N")
            break
        }

        pos = (pos + D + 16) % 16
    }
}