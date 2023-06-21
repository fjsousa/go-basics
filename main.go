package main

import "fmt"

func sum(a, b int) int {
    return a + b
}

func main() {
    a := 5
    b := 10
    result := sum(a, b)
    fmt.Printf("Sum of %d and %d is %d\n", a, b, result)
}
