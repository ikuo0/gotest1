
package kmeans

import (
	//"fmt"
	"math"
	"math/rand"
	//"os"
	//"rand"
	n1d "github.com/ikuo0/gotest1/num1d"
	n2d "github.com/ikuo0/gotest1/num2d"
	//"github.com/ikuo0/gotest1/numlib"
)
func InitRandom(n_clusters int, m n2d.Mat) (n1d.I64Arr, n2d.Mat) {
	rSize, _ := n2d.Size(m)
	meanIdxs := n1d.IntArange(0, rSize, 1)
	meanIdxs = n1d.IntShuffle(meanIdxs)
	meanIdxs = meanIdxs[:3]
	means := n2d.Indexing(m, meanIdxs)
	return meanIdxs, means
}

func DifNorm(a, b n1d.F64Arr) (float64) {
	_, dif := n1d.NNSubtract(a, b)
	return n1d.SumNorm(dif)
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
				distances[r] = DifNorm(center, m[r])
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

func Predict(n_clusters int, means n2d.Mat, x n1d.F64Arr) (int, float64) {
	norms := n1d.Create(n_clusters)
	for cluster := 0; cluster < n_clusters; cluster += 1 {
		norms[cluster] = DifNorm(x, means[cluster])
	}
	idx := n1d.ArgMin(norms)
	return idx, norms[idx]
}

func InitImprovisation(n_clusters int, x n2d.Mat) (n1d.I64Arr, n2d.Mat) {
	//n_clusters = n_clusters + 1
	rSize, cSize := n2d.Size(x)
	means := n2d.Zeros(n_clusters, cSize)
	idxs := n1d.IntZeros(n_clusters)
	var calcMinClusterDistance = func(_means n2d.Mat) (float64, int, int) {
		_minDistance := math.MaxFloat64
		_minDistanceIdx1 := 0
		_minDistanceIdx2 := 0
		for i := 0; i < n_clusters; i++ {
			for j := 0; j < i; j++ {
				if j == i {
					continue
				} else {
					distance := DifNorm(_means[i], _means[j])
					if distance < _minDistance {
						_minDistance = distance
						_minDistanceIdx1 = i
						_minDistanceIdx2 = j
					}
				}
			}
		}
		return _minDistance, _minDistanceIdx1, _minDistanceIdx2
	}
	minDistance := math.SmallestNonzeroFloat64
	minIdx1, minIdx2 := 0, 0
	for cnt := 0; cnt < 2; cnt ++ {
		for r := 0; r < rSize; r++ {
			ix := x[r]
			_, distance := Predict(n_clusters, means, ix)
			if distance > minDistance {
				norm1 := DifNorm(ix, means[minIdx1])
				norm2 := DifNorm(ix, means[minIdx2])
				if norm1 < norm2 {
					means[minIdx1] = ix
					idxs[minIdx1] = r
				} else {
					means[minIdx2] = ix
					idxs[minIdx2] = r
				}
				minDistance, minIdx1, minIdx2 = calcMinClusterDistance(means)
			} else {
				//means[predict] = ix
				//idxs[predict] = r
			}
		}
	}
	//fmt.Println(means)
	//os.Exit(9)
	//return idxs[1:], means[1:]
	return idxs, means
}

func Estep(n_clusters int, means n2d.Mat, x n2d.Mat) (n2d.Mat, n1d.F64Arr) {
	length, _ := n2d.Size(x)
	//norms := n2d.Create(length, n_clusters)
	predicts := n1d.IntZeros(length)
	distances := n1d.Zeros(length)
	for r := 0; r < length; r++ {
		predict, distance := Predict(n_clusters, means, x[r])
		predicts[r] = predict
		distances[r] = distance
		//for cluster := 0; cluster < n_clusters; cluster += 1 {
		//	_, dif := n1d.NNSubtract(x[r], means[cluster])
		//	norms[r][cluster] = n1d.SumNorm(dif)
		//}
	}
	probability := n2d.Zeros(length, n_clusters)
	for i, pred := range(predicts) {
		probability[i][pred] = 1.0
	}
	return probability, distances
}

func Mstep(n_clusters int, probability n2d.Mat, x n2d.Mat) (n2d.Mat) {
	rSize, cSize := n2d.Size(x)
	newMeans := n2d.Create(n_clusters, cSize)
	total := n2d.Opt().Axis(n2d.ConstAxisRow).Total(probability)
	for cluster := 0; cluster < n_clusters; cluster += 1 {
		sum := n1d.Zeros(cSize)
		for r := 0; r < rSize; r += 1 {
			ix := n1d.N1Multiple(x[r], float64(probability[r][cluster]))
			_, sum = n1d.NNPlus(sum, ix)
		}
		newMeans[cluster] = n1d.N1Division(sum, total[cluster])
	}
	return newMeans
}

func MeansShiftTotal(n_clusters int, a n2d.Mat, b n2d.Mat) (float64) {
	//center_shift_total = squared_norm(centers_old - centers)
	_, sub := n2d.NMSubtract(a, b)
	return n2d.SquaredNorm(sub)
}
