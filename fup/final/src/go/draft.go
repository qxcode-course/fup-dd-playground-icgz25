package main
import "fmt"
func main() {
    var nota1, nota2, notafinal int
    fmt.Scan(&nota1, &nota2, &notafinal)

    media := float64 (nota1+nota2) / 2.0

    if media >= 7.0 {
        fmt.Println("aprovado")
    } else if media < 4.0 {
        fmt.Println("reprovado")
    } else {
        novamedia := (media / 2.0)

        if novamedia >= 5.0 {
            fmt.Println("aprovado na final")
        } else {
            fmt.Println("reprovado na final")
        }
    }
}
