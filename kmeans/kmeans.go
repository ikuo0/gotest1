
package kmeans

import (
	//"fmt"
	"math/rand"
	//"rand"
	n1d "github.com/ikuo0/gotest1/num1d"
	n2d "github.com/ikuo0/gotest1/num2d"
	//"github.com/ikuo0/gotest1/numlib"
)
func InitKmeansRandom(n_clusters int, m n2d.Mat) (n1d.I64Arr, n2d.Mat) {
	rSize, _ := n2d.Size(m)
	meanIdxs := n1d.IntArange(0, rSize, 1)
	meanIdxs = n1d.IntShuffle(meanIdxs)
	meanIdxs = meanIdxs[:3]
	means := n2d.Indexing(m, meanIdxs)
	return meanIdxs, means
}

func InitKmeansPlusPlus(n_clusters int, m n2d.Mat) (n1d.I64Arr, n2d.Mat) {
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
			//fmt.Println("centerIdx", centerIdx)
		}
		meanIdxs[cluster] = centerIdx
	}
	means := n2d.Indexing(m, meanIdxs)
	return meanIdxs, means
}

func Estep(n_clusters int, means n2d.Mat, x n2d.Mat) (n2d.Mat) {
	length, _ := n2d.Size(x)
	norms := n2d.Create(length, n_clusters)
	for r := 0; r < length; r++ {
		for cluster := 0; cluster < n_clusters; cluster += 1 {
			_, dif := n1d.NNSubtract(x[r], means[cluster])
			norms[r][cluster] = n1d.SumNorm(dif)
		}
	}
	boolIndex := n2d.Zeros(n_clusters, length)
	for r := 0; r < length; r++ {
		cluster := n1d.ArgMin(norms[r])
		boolIndex[cluster][r] = 1
	}
	return boolIndex
}

func Mstep(n_clusters int, boolIndex n2d.Mat, x n2d.Mat) (n2d.Mat) {
	rSize, cSize := n2d.Size(x)
	newMeans := n2d.Create(n_clusters, cSize)
	for cluster := 0; cluster < n_clusters; cluster += 1 {
		sum := n1d.Zeros(cSize)
		for r := 0; r < rSize; r += 1 {
			ix := n1d.N1Multiple(x[r], float64(boolIndex[cluster][r]))
			_, sum = n1d.NNPlus(sum, ix)
		}
		subTotal := n1d.SumTotal(boolIndex[cluster])
		newMeans[cluster] = n1d.N1Division(sum, subTotal)
	}
	return newMeans
}

func MeansShiftTotal(n_clusters int, a n2d.Mat, b n2d.Mat) (float64) {
	//center_shift_total = squared_norm(centers_old - centers)
	_, sub := n2d.NMSubtract(a, b)
	return n2d.SquaredNorm(sub)
}
