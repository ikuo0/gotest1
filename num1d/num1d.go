package num1d

import (
    "math"
    "math/rand"
)

type TypeRandom int
const (
    ConstRandomUniform TypeRandom = iota
    ConstRandomNormal
)

type F64Arr []float64

func Create(size int) F64Arr {
    arr := make(F64Arr, size)
    return arr
}

func Zeros(size int) F64Arr {
    return Create(size)
}

func Ones(size int) F64Arr {
    arr := Zeros(size)
    for i := 0; i < len(arr); i++ {
        arr[i] = 1.0
    }
    return arr
}

func Arange(a float64, b float64, count int) F64Arr {
    arr := Zeros(int(count))
    step := (b - a) / float64(count-1)
    for i := 0; i < count; i++ {
        arr[i] = a + step*float64(i)
    }
    return arr
}

func Rand(size int, randomType TypeRandom) F64Arr {
    arr := Create(size)
    fn := rand.Float64
    if randomType == ConstRandomUniform {
        fn = rand.Float64
    } else if randomType == ConstRandomNormal {
        fn = rand.NormFloat64
    }
    for i := 0; i < size; i++ {
        arr[i] = fn()
    }
    return arr
}

type I64Arr []int

func IntCreate(size int) I64Arr {
    arr := make(I64Arr, size)
    return arr
}

func IntFull(size int, n int) I64Arr {
    arr := IntCreate(size)
    for i := 0; i < len(arr); i++ {
        arr[i] = n
    }
    return arr
}

func ToInt(arr F64Arr) []int {
    res := IntCreate(int(len(arr)))
    for i := 0; i < len(arr); i++ {
        res[i] = int(arr[i])
    }
    return res
}

func IntWhereEq(arr I64Arr, n int) I64Arr {
    idxs := []int{}
    for i := 0; i < len(arr); i++ {
        if arr[i] == n {
            idxs = append(idxs, i)
        }
    }
    return idxs
}

////////////////////////////////////////
// N:1 calculation
////////////////////////////////////////
func N1Plus(arr F64Arr, n float64) F64Arr {
    res := Create(int(len(arr)))
    for i := 0; i < len(arr); i++ {
        res[i] = arr[i] + n
    }
    return res
}

func N1Subtract(arr F64Arr, n float64) F64Arr {
    res := Create(int(len(arr)))
    for i := 0; i < len(arr); i++ {
        res[i] = arr[i] - n
    }
    return res
}

func N1Multiple(arr F64Arr, n float64) F64Arr {
    res := Create(int(len(arr)))
    for i := 0; i < len(arr); i++ {
        res[i] = arr[i] * n
    }
    return res
}

func N1Division(arr F64Arr, n float64) F64Arr {
    res := Create(int(len(arr)))
    for i := 0; i < len(arr); i++ {
        res[i] = arr[i] / n
    }
    return res
}

func N1Sqrt(arr F64Arr) F64Arr {
    res := Create(int(len(arr)))
    for i := 0; i < len(arr); i++ {
        res[i] = math.Sqrt(arr[i])
    }
    return res
}

func N1Pow(arr F64Arr, base float64) F64Arr {
    res := Create(int(len(arr)))
    for i := 0; i < len(arr); i++ {
        res[i] = math.Pow(arr[i], base)
    }
    return res
}

////////////////////////////////////////
// N:N calculation
////////////////////////////////////////
func NNPlus(a F64Arr, b F64Arr) (bool, F64Arr) {
    if len(a) != len(b) {
        return false, nil
    }
    res := Create(int(len(a)))
    for i := 0; i < len(a); i++ {
        res[i] = a[i] + b[i]
    }
    return true, res
}

func NNSubtract(a F64Arr, b F64Arr) (bool, F64Arr) {
    if len(a) != len(b) {
        return false, nil
    }
    res := Create(int(len(a)))
    for i := 0; i < len(a); i++ {
        res[i] = a[i] - b[i]
    }
    return true, res
}

func NNMulti(a F64Arr, b F64Arr) (bool, F64Arr) {
    if len(a) != len(b) {
        return false, nil
    }
    res := Create(int(len(a)))
    for i := 0; i < len(a); i++ {
        res[i] = a[i] * b[i]
    }
    return true, res
}

func NNDivision(a F64Arr, b F64Arr) (bool, F64Arr) {
    if len(a) != len(b) {
        return false, nil
    }
    res := Create(int(len(a)))
    for i := 0; i < len(a); i++ {
        res[i] = a[i] / b[i]
    }
    return true, res
}

////////////////////////////////////////
// Sum
////////////////////////////////////////
func SumTotal(arr F64Arr) float64 {
    res := float64(0)
    for i := 0; i < len(arr); i++ {
        res += arr[i]
    }
    return res
}
func SumMean(arr F64Arr) float64 {
    total := SumTotal(arr)
    return total / float64(len(arr))
}
func SumNorm(arr F64Arr) float64 {
    squared := N1Pow(arr, 2)
    total := SumTotal(squared)
    return math.Sqrt(total)
}


type Option struct {
    randomType TypeRandom
}

func (me* Option) RandomType(n TypeRandom) (*Option) {
    me.randomType = n
    return me
}

func (me* Option) Rand(size int) (F64Arr){
    return Rand(size, me.randomType)
}

func Opt() (*Option) {
    return &Option {
        randomType: ConstRandomUniform,
    }
}

