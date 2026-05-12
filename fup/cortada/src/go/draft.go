package main
import ("fmt")
func main() {
    var B, T int
    fmt.Scanln(&B)
    fmt.Scanln(&T)

    total := 160 * 70

    areaEsquerda := (B + T) * 70 / 2
    areaDireita := total - areaEsquerda

    if areaEsquerda > total/2 {
        fmt.Println(1)
    } else if areaDireita > total/2 {
        fmt.Println(2)
    } else {
        fmt.Println(0)
    }
}