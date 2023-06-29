package reversestr

func Rev(s string) string {
	runes := []rune(s)
	lenght := len(runes)

	for i, j := 0, lenght - 1; i<j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}