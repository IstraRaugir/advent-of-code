package utils

func SumIntSlice(section []int) (output int) {
	for _, value := range section {
		output += value
	}
	return
}

func RuneSliceToIntSlice(commonItems []rune) {
	var intSlice []int
	for _, v := range commonItems {
		intSlice = append(intSlice, int(v))
	}
}
