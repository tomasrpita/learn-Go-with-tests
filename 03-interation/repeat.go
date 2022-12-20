package interation

// Repeat returns a string with a character n interations rpeated
func Repeat(character string, iterations int) string {
	var repeated = ""
	for i := 0; i < iterations; i++ {
		repeated += character
	}
	return repeated
}
