package histogram

func ComputeCDF(hist [256]int) [256]int {
	var cdf [256]int
	cdf[0] = hist[0]
	for i := 1; i < 256; i++ {
		cdf[i] = cdf[i-1] + hist[i]
	}
	return cdf
}
