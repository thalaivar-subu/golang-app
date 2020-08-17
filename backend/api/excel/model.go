package excel

func ExcelRowNames(n int, input string) []string {
	result := []string{input}
	tempString := result[0]
	for i := 0; i < n; i++ {
		tempString = iterateString(tempString)
		result = append(result, tempString)
	}
	return result
}

func iterateString(x string) string {
	runeSlice := []rune(x)
	charArray := make([]int, 0)

	for i := 0; i < len(runeSlice); i++ {
		charArray = append(charArray, int(runeSlice[i])%64)
	}
	arrayLength := len(charArray)
	for i := arrayLength - 1; i >= 0; i-- {
		if charArray[i] == 26 {
			charArray[i] = 1
			if !((i-1) < len(charArray) && ((i - 1) >= 0)) {
				charArray = append([]int{0}, charArray...)
				arrayLength = arrayLength + 1
				i = i + 1
			}
		} else {
			charArray[i] = charArray[i] + 1
			break
		}
	}
	result := ""
	for i := 0; i < len(charArray); i++ {
		result += string(charArray[i] + 64)
	}
	return string(result)
}
