
package kmeans

import (
	"fmt"
	"math/rand"
	//"rand"
	n1d "github.com/ikuo0/gotest1/num1d"
	n2d "github.com/ikuo0/gotest1/num2d"
	//"github.com/ikuo0/gotest1/numlib"
)

func InitKmeansPlusPlus(m n2d.Mat, n_clusters int) {
	rSize, _ := n2d.Size(m)
	meanIdxs := n1d.IntFull(n_clusters, -1)
	for cluster := 0; cluster < n_clusters; cluster++ {
		centerIdx := 0
		if cluster == 0 {
			centerIdx = rand.Intn(rSize)
		} else {
			center := m[centerIdx]
			distances := n1d.Zeros(rSize)
			for r := 0; r < rSize; r++ {
				_, dif := n1d.NNSubtract(center, m[r])
				distances[r] = n1d.SumNorm(dif)
			}
			pow := n1d.N1Pow(distances, 2)
			denom := n1d.SumTotal(pow)
			probability := float64(0)
			probabilities := n1d.Zeros(rSize)
			for r := 0; r < rSize; r++ {
				probability += distances[r] / denom
				probabilities[r] = probability
			}
			//fmt.Println(probabilities)
			for {
				pos := rand.Float64() * probabilities[len(probabilities) - 1]
				for r := rSize - 1; r >= 0; r-- {
					if probabilities[r] <= pos {
						centerIdx = r
						break
					}
				}
				idxs := n1d.IntWhereEq(meanIdxs, centerIdx)
				if len(idxs) == 0 {
					break
				}
			}
			//fmt.Println("pos", pos)
			fmt.Println("centerIdx", centerIdx)
		}
		meanIdxs[cluster] = centerIdx
	}
}

