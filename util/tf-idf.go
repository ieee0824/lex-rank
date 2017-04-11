package util

func CalcTFIDF(w string, texts []string) float64 {
	return CalcIDF(w, texts) * CalcTF(w)
}
