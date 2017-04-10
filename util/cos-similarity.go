package util

import (
	"math"
	"sort"
)

func calcIndex(v []string) map[string]int {
	ret := map[string]int{}
	for _, w := range v {
		_, ok := ret[w]
		if ok {
			ret[w]++
		} else {
			ret[w] = 1
		}
	}
	return ret
}

func aggregateIndex(a, b map[string]int) map[string]int {

	var ret = map[string]int{}
	for k, v := range a {
		ret[k] = v
	}
	for k, v := range b {
		_, ok := b[k]
		if ok {
			ret[k] = b[k] + v
		} else {
			ret[k] = v
		}
	}
	return ret
}

func calcVec(ownIndex, index map[string]int) []int {
	ret := []int{}
	keys := make([]string, len(index))

	i := 0
	for k, _ := range index {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for _, k := range keys {
		n, ok := ownIndex[k]
		if ok {
			ret = append(ret, n)
		} else {
			ret = append(ret, 0)
		}
	}
	return ret
}

func mulVec(a, b []int) int64 {
	var ret int64

	for i, _ := range a {
		ret = ret + int64(a[i]*b[i])
	}
	return ret
}

func lenVec(vec []int) float64 {
	var buffer int64

	for _, v := range vec {
		buffer = buffer + int64(v*v)
	}

	return math.Sqrt(float64(buffer))
}

func CalcCosSimil(A []string, B []string) float64 {
	index := map[string]int{}
	finA := make(chan map[string]int)
	finB := make(chan map[string]int)

	go func(ch chan map[string]int) {
		finA <- calcIndex(A)
	}(finA)
	go func(ch chan map[string]int) {
		finB <- calcIndex(B)
	}(finB)

	indexA := <-finA
	indexB := <-finB

	index = aggregateIndex(indexA, indexB)

	vecA := calcVec(indexA, index)
	vecB := calcVec(indexB, index)

	AB := mulVec(vecA, vecB)
	lenA := lenVec(vecA)
	lenB := lenVec(vecB)

	return float64(AB) / (lenA * lenB)
}
