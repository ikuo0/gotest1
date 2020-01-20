package num1d

import (
	"math"
)

//const Const

type F64Arr []float64

func Create(size uint) F64Arr {
	arr := make(F64Arr, size)
	return arr
}

func Zeros(size uint) F64Arr {
	return Create(size)
}

func Ones(size uint) F64Arr {
	arr := Zeros(size)
	for i := 0; i < len(arr); i++ {
		arr[i] = 1.0
	}
	return arr
}

func Arange(a float64, b float64, count int) F64Arr {
	arr := Zeros(uint(count))
	step := (b - a) / float64(count-1)
	for i := 0; i < count; i++ {
		arr[i] = a + step*float64(i)
	}
	return arr
}

type I64Arr []int

func CreateInt(size uint) I64Arr {
	arr := make(I64Arr, size)
	return arr
}

func ToInt(arr F64Arr) []int {
	res := CreateInt(uint(len(arr)))
	for i := 0; i < len(arr); i++ {
		res[i] = int(arr[i])
	}
	return res
}

////////////////////////////////////////
// N:1 calculation
////////////////////////////////////////
func N1Plus(arr F64Arr, n float64) F64Arr {
	res := Create(uint(len(arr)))
	for i := 0; i < len(arr); i++ {
		res[i] = arr[i] + n
	}
	return res
}

func N1Subtract(arr F64Arr, n float64) F64Arr {
	res := Create(uint(len(arr)))
	for i := 0; i < len(arr); i++ {
		res[i] = arr[i] - n
	}
	return res
}

func N1Multiple(arr F64Arr, n float64) F64Arr {
	res := Create(uint(len(arr)))
	for i := 0; i < len(arr); i++ {
		res[i] = arr[i] * n
	}
	return res
}

func N1Division(arr F64Arr, n float64) F64Arr {
	res := Create(uint(len(arr)))
	for i := 0; i < len(arr); i++ {
		res[i] = arr[i] / n
	}
	return res
}

func N1Sqrt(arr F64Arr) F64Arr {
	res := Create(uint(len(arr)))
	for i := 0; i < len(arr); i++ {
		res[i] = math.Sqrt(arr[i])
	}
	return res
}

func N1Pow(arr F64Arr, base float64) F64Arr {
	res := Create(uint(len(arr)))
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
	res := Create(uint(len(a)))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] + b[i]
	}
	return true, res
}

func NNSubtract(a F64Arr, b F64Arr) (bool, F64Arr) {
	if len(a) != len(b) {
		return false, nil
	}
	res := Create(uint(len(a)))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] - b[i]
	}
	return true, res
}

func NNMulti(a F64Arr, b F64Arr) (bool, F64Arr) {
	if len(a) != len(b) {
		return false, nil
	}
	res := Create(uint(len(a)))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] * b[i]
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
