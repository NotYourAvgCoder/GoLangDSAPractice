package strings

func LengthOfLongestSubstring(s string) int {

	greatestLength := 0
	s_runes := []rune(s)

	for i := 0; i < len(s_runes); i++ {
		hashMap := make(map[rune]bool)
		currentLength := 0
		for j := i; j < len(s_runes); j++ {
			if hashMap[s_runes[j]] {
				hashMap = make(map[rune]bool)
				if currentLength > greatestLength {
					greatestLength = currentLength
				}
				currentLength = 0
			}
			currentLength += 1
			hashMap[s_runes[j]] = true
		}
		if currentLength > greatestLength {
			greatestLength = currentLength
		}
	}

	return greatestLength
}
