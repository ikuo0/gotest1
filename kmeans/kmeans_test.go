
package kmeans

import (
	"testing"
	"fmt"
	"log"
	"math/rand"
	"time"
	//n1d "github.com/ikuo0/gotest1/num1d"
	n2d "github.com/ikuo0/gotest1/num2d"
	"github.com/ikuo0/gotest1/numlib"
)

func TestKmeansPlusPlus(t* testing.T) {
	rand.Seed(time.Now().UnixNano())
	if ok, m := numlib.LoadTxt("./iris.txt"); ok != true {
		log.Fatal("load iris error")
	} else {
		m = n2d.SelectColumn(m, []int{0, 1, 2, 3})
		mean, std := numlib.StandardScalerFit(m)
		m = numlib.StandardScalerTransform(mean, std, m)
		fmt.Println(m[0])
		InitKmeansPlusPlus(m, 3)
	}
}
