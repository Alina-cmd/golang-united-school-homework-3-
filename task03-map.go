package homework

func sortMapValues(input map[int]string) (result []string) {
	var minn int
	for key, _ := range input {
		minn = key
		break
	}
	i := 0
	length := len(input)
	for i < length {
		for key, _ := range input {
			if minn > key {
				minn = key
			}
		}
		result = append(result, input[minn])
		i++
		delete(input, minn)
		for key, _ := range input {
			minn = key
			break
		}
	}
	return
}
