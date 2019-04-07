package utility

func SliceAverageFloat(xs []float64) float64 {
	if len(xs) == 0 {
		return 0.0
	}

	total := 0.0
	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}

func SliceSumFloat(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total
}

func SliceStdFloat(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total
}

func MaxInt(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func MinInt(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func SliceMaxFloat(xs []float64) float64 {
	if len(xs) == 0 {
		return 0.0
	}

	max := 0.0
	for _, v := range xs {
		if v > max {
			max = v
		}
	}

	return max
}
