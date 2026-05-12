package main
import "fmt"
func main() {
    var jog1, jog2, jog3 int

    fmt.Scanln(&jog1)
    fmt.Scanln(&jog2)
    fmt.Scanln(&jog3)

    if jog1 != jog2 && jog1 != jog3 {
        fmt.Println("jog1")
    } else if jog2 != jog1 && jog2 != jog3 {
        fmt.Println("jog2")
    } else if jog3 != jog1 && jog3 != jog2 {
        fmt.Println("jog3")
    } else {
        fmt.Println("empate")
    }
}