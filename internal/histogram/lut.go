package histogram

import "github.com/kylin419/CVforge/internal/kernel"

func MinNonZeroCDF(cdf [256]int) int {
	for _, value := range cdf {
		if value != 0 {
			return value
		}
	}
	return 0
}

func BuildLUT(cdf [256]int, total int) [256]uint8 {
	min_cdf := MinNonZeroCDF(cdf)
	var result [256]uint8
	if total == min_cdf {
		return result
	}
	for i := 0; i < len(cdf); i++ {
		new := (float64(cdf[i]-min_cdf) / float64(total-min_cdf)) * 255.0
		result[i] = kernel.Clamp(new)
	}
	return result
}
