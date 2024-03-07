package pugutils

func Lerp64(valA, valB, t float64) float64 {
	return valA + (valB-valA)*t
}


func Lerp32(valA, valB, t float32) float32 {
	return valA + (valB-valA)*t
}


