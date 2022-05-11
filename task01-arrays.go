package homework

func average(input [15]float32) (result float32) {
	var sum float32
	for _, val := range input {
		sum += val
	}
	result = sum / float32(len(input))
	return
}
