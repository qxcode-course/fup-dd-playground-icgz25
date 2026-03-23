package main
import "fmt"
func main() {
    var q1, q2, q3 float64
    var p1, p2, p3 float64
    var dinheiro float64

    fmt.Scan(&q1)
    fmt.Scan(&q2)
    fmt.Scan(&q3)
    fmt.Scan(&p1)
    fmt.Scan(&p2)
    fmt.Scan(&p3)
    fmt.Scan(&dinheiro)

    total := q1*p1 + q2*p2 + q3*p3
    troco := dinheiro - total

    fmt.Printf("%.2f\n", troco)
}
