package util

import (
	"math"
	"strings"
)

func CalcIDF(w string, texts []string) float64 {
	var d int
	var D = len(texts)
	for _, t := range texts {
		if strings.Contains(t, w) {
			d++
		}
	}

	return -1 * math.Log(float64(d)/float64(D))
}
