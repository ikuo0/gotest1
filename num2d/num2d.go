
package num2d

import (
    //"fmt"
    "math"
    //"os"
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

func Zeros(r int, c int) (Mat) {
    res := make([]n1d.F64Arr, r)
    for i := 0; i < r; i++ {
        res[i] = n1d.Zeros(c)
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

func Size(m Mat) (int, int) {
    return len(m), len(m[0])
}

////////////////////////////////////////
// etc
////////////////////////////////////////
func ArgMax(m Mat, axis TypeAxis) (n1d.I64Arr) {
    if axis != ConstAxisRow {
        m = Transpose(m)
    }
    rSize, _ := Size(m)
    res := n1d.IntZeros(rSize)
    for r := 0; r < rSize; r++ {
        res[r] = n1d.ArgMax(m[r])
    }
    return res
}

////////////////////////////////////////
// formation
////////////////////////////////////////
func Transpose(m Mat) (Mat) {
    rSize, cSize := Size(m)
    res := Create(cSize, rSize)
    for r := 0; r < rSize; r++ {
        for c := 0; c < cSize; c++ {
            res[c][r] = m[r][c]
        }
    }
    return res
}

func SelectColumn(m Mat, columns []int) (Mat) {
    rSize, _ := Size(m)
    res := Create(rSize, len(columns))
    for r := 0; r < rSize; r++ {
        for dstCol, srcCol := range(columns) {
            res[r][dstCol] = m[r][srcCol]
        }
    }
    return res
}

func Indexing(m Mat, idxs n1d.I64Arr) (Mat) {
    _, cSize := Size(m)
    res := Create(len(idxs), cSize)
    for i, idx := range(idxs) {
        res[i] = m[idx]
    }
    return res
}

func Flatten(m Mat) (n1d.F64Arr) {
    rSize, cSize := Size(m)
    res := n1d.Create(rSize * cSize)
    for r := 0; r < rSize; r++ {
        for c := 0; c < cSize; c++ {
            res[r * cSize + c] = m[r][c]
        }
    }
    return res
}


////////////////////////////////////////
// N x M calc
////////////////////////////////////////
func EqualyCheck(a Mat, b Mat) (bool) {
    rSizeA, cSizeA := Size(a)
    rSizeB, cSizeB := Size(b)
    return rSizeA == rSizeB && cSizeA == cSizeB
}
func NMSubtract(a Mat, b Mat) (bool, Mat) {

    if EqualyCheck(a, b) == false {
        return false, nil
    } else {
        rSize, cSize := Size(a)
        m := Create(rSize, cSize)
        for r := 0; r < rSize; r++ {
            for c := 0; c < cSize; c++ {
                m[r][c] = a[r][c] - b[r][c]
            }
        }
        return true, m
    }
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
    if axis == ConstAxisRow {
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

func SquaredNorm(x Mat) (float64) {
    rSize, cSize := Size(x)
    var total float64 = 0
    for r := 0; r < rSize; r++ {
        for c := 0; c < cSize; c++ {
            total += math.Pow(x[r][c], 2)
        }
    }
    return total
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

func (me *Option) Total(m Mat) (n1d.F64Arr) {
    return Total(m, me.axisType)
}

func (me *Option) Mean(m Mat) (n1d.F64Arr) {
    return Mean(m, me.axisType)
}

func (me *Option) Std(m Mat) (n1d.F64Arr) {
    return Std(m, me.axisType)
}

func (me *Option) ArgMax(m Mat) (n1d.I64Arr) {
    return ArgMax(m, me.axisType)
}


func Opt() (*Option) {
    return &Option{
        axisType: ConstAxisRow,
        randomType: ConstRandomUniform,
    }
}
