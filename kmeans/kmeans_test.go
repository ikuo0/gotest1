
package kmeans

import (
	"testing"
	"fmt"
	"log"
	"math/rand"
	//"os"
	"time"
	//n1d "github.com/ikuo0/gotest1/num1d"
	n2d "github.com/ikuo0/gotest1/num2d"
	"github.com/ikuo0/gotest1/numlib"
)

func TestInitRandom(t* testing.T) {
	rand.Seed(time.Now().UnixNano())
	if ok, m := numlib.LoadTxt("./iris.txt"); ok != true {
		log.Fatal("load iris error")
	} else {
		m = n2d.SelectColumn(m, []int{0, 1, 2, 3})
		mean, std := numlib.StandardScalerFit(m)
		m = numlib.StandardScalerTransform(mean, std, m)
		fmt.Println("Init Random")
		for i := 0; i < 1000; i++ {
			idxs, _ := InitKmeansRandom(3, m)
			fmt.Println(idxs)
		}
	}
}

func TestInitPlusPlus(t* testing.T) {
	rand.Seed(time.Now().UnixNano())
	if ok, m := numlib.LoadTxt("./iris.txt"); ok != true {
		log.Fatal("load iris error")
	} else {
		m = n2d.SelectColumn(m, []int{0, 1, 2, 3})
		mean, std := numlib.StandardScalerFit(m)
		m = numlib.StandardScalerTransform(mean, std, m)
		fmt.Println("Init PlusPlus")
		for i := 0; i < 1000; i++ {
			idxs, _ := InitKmeansPlusPlus(3, m)
			fmt.Println(idxs)
		}
	}
}

func TestKmeans(t* testing.T) {
	nClusters := 3
	rand.Seed(time.Now().UnixNano())
	if ok, m := numlib.LoadTxt("./iris.txt"); ok != true {
		log.Fatal("load iris error")
	} else {
		m = n2d.SelectColumn(m, []int{0, 1, 2, 3})
		mean, std := numlib.StandardScalerFit(m)
		m = numlib.StandardScalerTransform(mean, std, m)
		_, means := InitKmeansPlusPlus(nClusters, m)
		tol := 1e-5
		for iter := 0; iter < 100; iter++ {
			boolIndex := Estep(nClusters, means, m)
			//fmt.Println(boolIndex)
			newMeans := Mstep(nClusters, boolIndex, m)
			//fmt.Println(newMeans)
			shiftTotal := MeansShiftTotal(nClusters, means, newMeans)
			fmt.Println("shiftTotal", shiftTotal)
			means = newMeans

			if shiftTotal < tol{
				fmt.Println("converged")
				break
			}
		}
	}
}
