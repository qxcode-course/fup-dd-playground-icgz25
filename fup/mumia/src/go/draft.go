package main
import "fmt"
func main() {
    var pessoa string
    var idade int
    fmt.Scan(&pessoa, &idade)

    if idade <= 12 {
        fmt.Println("crianca")
    } else if idade >= 18 {
        fmt.Println("adolescente")
    } else if idade <= 65 {
        fmt.Println("adulto")
    } else if idade <= 1000 {
        fmt.Println("idoso")
    } else {
        fmt.Println("mumia")
    }
    
    
    
}
