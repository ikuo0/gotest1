
package numlib

import (
	"fmt"
	"log"
	"testing"
	n2d "github.com/ikuo0/gotest1/num2d"
)

func Test1(t* testing.T) {
	m := n2d.Opt().RandomType(n2d.ConstRandomUniform).Rand(100, 4)
	SaveTxt("./test.txt", m)
}

func Test2(t* testing.T) {
	if ok, m := LoadTxt("./test.txt"); ok == false {
		log.Fatal("LoadTxt error")
	} else {
		fmt.Println(m)
	}
}

