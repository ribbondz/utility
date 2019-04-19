package utility

func SliceCopyFloat(list []float64) (dst []float64) {
	dst = make([]float64, len(list))
	copy(dst, list)
	return
}

func SliceCopyFloatFloat(list [][]float64) [][]float64 {
	dst := make([][]float64, len(list))
	for i, v := range list {
		dst[i] = SliceCopyFloat(v)
	}
	return dst
}

func SliceCopyInt(list []int) (result []int) {
	result = make([]int, len(list))
	copy(result, list)
	return
}

func SliceCopyIntInt(list [][]int) [][]int {
	dst := make([][]int, len(list))
	for i, v := range list {
		dst[i] = SliceCopyInt(v)
	}
	return dst
}

func SliceCopyString(list []string) (result []string) {
	result = make([]string, len(list))
	copy(result, list)
	return
}

func MapCopyIntInt(oldMap map[int]int) map[int]int {
	result := make(map[int]int)
	for k, v := range oldMap {
		result[k] = v
	}
	return result
}

func MapCopyIntFloat(oldMap map[int]float64) map[int]float64 {
	result := make(map[int]float64)
	for k, v := range oldMap {
		result[k] = v
	}
	return result
}

func MapCopyIntString(oldMap map[int]string) map[int]string {
	result := make(map[int]string)
	for k, v := range oldMap {
		result[k] = v
	}
	return result
}

func MapCopyFloatInt(oldMap map[float64]int) map[float64]int {
	result := make(map[float64]int)
	for k, v := range oldMap {
		result[k] = v
	}
	return result
}

func MapCopyFloatFloat(oldMap map[float64]float64) map[float64]float64 {
	result := make(map[float64]float64)
	for k, v := range oldMap {
		result[k] = v
	}
	return result
}

func MapCopyFloatString(oldMap map[float64]string) map[float64]string {
	result := make(map[float64]string)
	for k, v := range oldMap {
		result[k] = v
	}
	return result
}
