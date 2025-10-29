package math

// Because after 10 years golang is still refusing to provide basic stdlib functions
func Abs(v float32) float32 {
	if v > 0 {
		return v
	}
	return v * -1
}

// Because after 10 years golang is still refusing to provide basic stdlib functions
func Min(a float32, b float32) float32 {
	if a > b {
		return b
	}
	return a
}
