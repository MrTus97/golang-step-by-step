package helper

func SumFromZero(number int) int {
	temp := 0
	for i := 0; i < number; i++ {
		temp += i
	}

	return temp
}
