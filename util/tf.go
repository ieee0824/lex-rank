package util

func CalcTF(w []string) map[string]float64 {
	ret := map[string]float64{}
	buffer := map[string]int{}

	for _, s := range w {
		_, ok := buffer[s]
		if ok {
			buffer[s]++
		} else {
			buffer[s] = 1
		}
	}

	N := len(buffer)
	for k, n := range buffer {
		ret[k] = float64(n) / float64(N)
	}

	return ret
}
