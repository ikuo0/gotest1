package num1d

import (
    "fmt"
    "testing"
)

func TestCreate(t *testing.T) {
    // Zeros
    a1 := Zeros(10)
    fmt.Println(a1)

    // Ones
    a2 := Ones(11)
    fmt.Println(a2)

    // Arange
    a3 := Arange(1, 3, 13)
    fmt.Println(a3)

    // ToInt
    a4 := ToInt(a3)
    fmt.Println(a4)
}

func TestN1Calc(t *testing.T) {
    a1 := Ones(10)
    fmt.Println(a1)
    a2 := N1Plus(a1, 1)
    fmt.Println(a2)
    a3 := N1Subtract(a1, 1)
    fmt.Println(a3)
    a4 := N1Multiple(a1, 2)
    fmt.Println(a4)
    a5 := N1Division(a1, 3)
    fmt.Println(a5)
}

func TestNNCalc(t *testing.T) {
    // Zeros
    a1 := Arange(0, 10, 11)
    a2 := Arange(1.1, 13.9, 11)
    fmt.Println(a1)
    fmt.Println(a2)

    // SumMean
    if ok, a3 := NNPlus(a1, a2); ok {
        fmt.Println(a3)
        fmt.Println(SumMean(a3))
    } else {
        fmt.Println("Error!")
    }

    // SumNorm
    a4 := SumNorm(a2)
    fmt.Println(a4)
}
