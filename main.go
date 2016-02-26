package main

import (
    "fmt"
    "math/rand"
)


const (
    rangex = 10.
    rangey = 10.
    steps  = 0
    tests  = 10
)


func equation(x []float64) []float64 {
    return []float64{x[0] * x[0] - x[1] * x[1]}
}


func genInput() []float64 {
    return []float64{rand.Float64() * rangex * 2 - rangex, rand.Float64() * rangey * 2 - rangey}
}

func main() {
    net := NewNetwork(0.2, 2, 10, 1)

    //  train it
    for i := 0; i < steps; i++ {
        x := genInput()
        ans := equation(x)
        fmt.Println("Learning",i,":", x[0], ",", x[1], "=", ans)
        net.Learn(x, equation(x))
    }

    //  test it
    for i := 0; i < tests; i++ {
        x := genInput()
        ans := equation(x)
        guess := net.Guess(x)
        fmt.Println("Guessing", x, "=", guess, ". Is", ans)
    }
}
