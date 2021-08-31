package math

func Average(nums []float64) float64 {
	total := float64(0)
	for _, v := range nums {
		total += v
	}
	return total / float64(len(nums))
}
