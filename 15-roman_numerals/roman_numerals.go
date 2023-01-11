package numerals

import "strings"

func ConvertToRoman(arabic int) string {

	var result strings.Builder
	// for i := 0; i < arabic; i++ {
	// 	if arabic == 5 {
	// 		result.WriteString("V")
	// 		break
	// 	}
	// 	if arabic == 4 {
	// 		result.WriteString("IV")
	// 		break
	// 	}

	// 	result.WriteString("I")
	// }

	for arabic > 0 {
		switch {
		case arabic > 8:
			result.WriteString("IX")
			arabic -= 9
		case arabic > 4:
			result.WriteString("V")
			arabic -= 5
		case arabic > 3:
			result.WriteString("IV")
			arabic -= 4
		default:
			result.WriteString("I")
			arabic--
		}
	}

	return result.String()

}
