package num2d

import (
	"testing"
	"fmt"
)

func Test1(t* testing.T) {
	m := Opt().RandomType(ConstRandomUniform).Rand(10, 3)
	//fmt.Println(m)
	mean := Opt().Axis(0).Mean(m)
	fmt.Println(mean)
	std := Opt().Axis(1).Std(m)
	fmt.Println(std)
}
