package validators

func PercentIsValid(percent float64) bool {
	return percent >= 0 && percent <= 100
}
