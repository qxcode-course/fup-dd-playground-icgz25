package main
import "fmt"
func main() {
    var a, b int
    
    fmt.Scan(&a, &b)

    resto := a % b
    quociente := a / b

    fmt.Println(quociente, resto)
}
