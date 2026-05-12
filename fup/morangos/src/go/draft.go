package main
import "fmt"
func main() {
    var c1, l1, c2, l2 int

    fmt.Scanln(&c1)
    fmt.Scanln(&l1)
    fmt.Scanln(&c2)
    fmt.Scanln(&l2)

    area1 := c1 * l1
    area2 := c2 * l2

    if area1 > area2 {
        fmt.Println(area1)
    } else {
        fmt.Println(area2)
    }
}
