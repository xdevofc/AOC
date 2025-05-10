package main

func BetweenLetters(line string) bool {

	for i := 2; i < len(line); i++ {
		if line[i-2] == line[i] {
			return true
		}
	}
	return false
	// aba  fef
}
