package main

import "fmt"

func main() {
    n := 50
    a, b := 0, 1

    for i := 0; i < n; i++ {
        fmt.Print(a, " ")
        a, b = b, a+b
    }
    fmt.Println()
}
