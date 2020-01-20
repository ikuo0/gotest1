package num1d

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

////////////////////////////////////////
// N:N calculation
////////////////////////////////////////
func NNPlus(a F64Arr, b F64Arr) F64Arr {

	res := Create()
}
