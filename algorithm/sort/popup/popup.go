package popup

func Run(nums *[]int) {
	length := len(*nums)
	if length < 2 {
		return
	}

	for i := 0; i < length-1; i++ {
		for j := i+1; j < length; j++ {
			if (*nums)[i] > (*nums)[j] {
				(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
			}
		}
	}
}