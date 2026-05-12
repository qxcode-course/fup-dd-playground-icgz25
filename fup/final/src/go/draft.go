package main
import "fmt"
func main() {
    var nota1, nota2, notaFinal int

    fmt.Scanln(&nota1)
    fmt.Scanln(&nota2)
    fmt.Scanln(&notaFinal)

    media := float64(nota1+nota2) / 2.0

    if media >= 7 {
        fmt.Println("aprovado")
    } else if media < 4 {
        fmt.Println("reprovado")
    } else {
        mediaFinal := (media + float64(notaFinal)) / 2.0
        if mediaFinal >= 5 {
            fmt.Println("aprovado na final")
        } else {
            fmt.Println("reprovado na final")
        }
    }
}
