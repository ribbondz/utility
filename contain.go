package utility

func FloatSliceContain(list []float64, value float64) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func IntSliceContain(list []int, value int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func StringSliceContain(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

