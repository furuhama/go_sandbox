package mytypes

// MyFloat is original float64 type
type MyFloat float64

// Abs is for MyFloat type method
func (f MyFloat) Abs() float64 {
	if f < 0 {
		f = f * -1
	}
	return float64(f)
}
