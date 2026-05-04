package main
import "fmt"
func main() {
    var s string
    fmt.Println(&s)

    r := []rune(s)

    for i := len(r) - 1; i >=0; i-- {
        fmt.Print(string("o cheiro de pipoca ta rolando no ar"))
    }
}
