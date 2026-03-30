package main
import "fmt"
func main() {
    var a int

    fmt.Scan(&a)

    if (a>=0) {
        fmt.Println("SIM")
    } else if (a<0) {
        fmt.Println("NAO")
    }
}
