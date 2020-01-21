
package num2d

import (
    n1d "github.com/ikuo0/gotest1/num1d"
)
//type F64Arr []float64
type Mat []n1d.F64Arr

const (
	ConstRandomUniform n1d.TypeRandom = n1d.ConstRandomUniform
	ConstRandomNormal = n1d.ConstRandomNormal
)

type TypeAxis int
const (
	ConstAxisRow TypeAxis = iota
	ConstAxisCol
)

////////////////////////////////////////
// create
////////////////////////////////////////
func Create(r int, c int) (Mat) {
	res := make([]n1d.F64Arr, r)
	for i := 0; i < r; i++ {
		res[i] = make([]float64, c)
	}
	return Mat(res)
}

func Rand(r int, c int, randomType n1d.TypeRandom) (Mat) {
	m := Create(r, c)
	for i := 0; i < r; i++ {
		m[i] = n1d.Opt().RandomType(randomType).Rand(c)
	}
	return m
}

func Transpose(m Mat) (Mat) {
    rowSize := len(m)
    colSize := len(m[0])
	res := Create(colSize, rowSize)
	for r := 0; r < rowSize; r++ {
		for c := 0; c < colSize; c++ {
			res[c][r] = m[r][c]
		}
	}
	return res
}

////////////////////////////////////////
// sum api
////////////////////////////////////////
func Total(m Mat, axis TypeAxis) (n1d.F64Arr) {
	if axis != ConstAxisRow {
		m = Transpose(m)
	}
    rowSize := len(m)
	colSize := len(m[0])
	res := n1d.Zeros(colSize)
	for r := 0; r < rowSize; r++ {
		for c := 0; c < colSize; c++ {
			res[c] += m[r][c]
		}
	}
	return res
}

func Mean(m Mat, axis TypeAxis) (n1d.F64Arr) {
    rowSize := len(m)
	colSize := len(m[0])
	res := Total(m, axis)
    if axis == 0 {
		return n1d.N1Division(res, float64(rowSize))
    } else {
		return n1d.N1Division(res, float64(colSize))
    }
}

func Std(m Mat, axis TypeAxis) (n1d.F64Arr) {
	if axis != ConstAxisRow {
		m = Transpose(m)
	}
    rowSize := len(m)
	colSize := len(m[0])
	mean := Mean(m, ConstAxisRow)
	res := n1d.Zeros(colSize)
	for r := 0; r < rowSize; r++ {
		_, a := n1d.NNSubtract(m[r], mean)
		a = n1d.N1Pow(a, 2)
		for c := 0; c < colSize; c++ {
			res[c] += a[c]
		}
	}
	res = n1d.N1Division(res, float64(rowSize))
	res = n1d.N1Sqrt(res)
	return res
}

////////////////////////////////////////
// Option instance
////////////////////////////////////////
type Option struct {
	axisType TypeAxis
	randomType n1d.TypeRandom
}

////////////////////////////////////////
// specify options
////////////////////////////////////////
func (me *Option) Axis(n TypeAxis) (*Option) {
    me.axisType = n
    return me
}

func (me *Option) RandomType(n n1d.TypeRandom) (*Option) {
    me.randomType = n
    return me
}

////////////////////////////////////////
// sum members
////////////////////////////////////////
func (me *Option) Rand(r int, c int) (Mat) {
	return Rand(r, c, me.randomType)
}

func (me *Option) Mean(m Mat) (n1d.F64Arr) {
	return Mean(m, me.axisType)
}

func (me *Option) Std(m Mat) (n1d.F64Arr) {
	return Std(m, me.axisType)
}


func Opt() (*Option) {
	return &Option{
		axisType: ConstAxisRow,
		randomType: ConstRandomUniform,
	}
}
