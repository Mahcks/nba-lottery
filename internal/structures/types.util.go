package structures

import (
	"hash/maphash"
	"math"
	"math/rand"
)

type Pair struct {
	Key        string
	Value      int
	FirstPicks int // times team got first pick
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].FirstPicks > p[j].FirstPicks }

func RandInt() int64 {
	r := rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))
	return r.Int63()
}

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}
