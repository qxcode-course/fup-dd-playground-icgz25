package main
import "fmt"
func main() {
    var wifi, login, admin int

    fmt.Scan(&wifi, &login, &admin)

    if (wifi == 0) {
        fmt.Println("you must connect to wifi")
        return
    }
    if login == 0 {
        fmt.Println("you need to login first")
        return
    }
    if admin == 0 {
        fmt.Println("you must to login as admin")
        return
    }

    fmt.Println("done")
}
