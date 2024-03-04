package pugutils

func Lerp(valA, valB, t float64) float64 {
	return valA + (valB-valA)*t
}


