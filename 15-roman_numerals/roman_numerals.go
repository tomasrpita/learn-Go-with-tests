package numerals

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (rn RomanNumerals) ValueOf(potentialNumber string) int {

	for i := 0; i < len(rn); i++ {
		if potentialNumber == rn[i].Symbol {
			return rn[i].Value
		}
	}

	return 0
}

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()

}

func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i] // it is a byte

		// look ahead to next symbol if we can and, the current symbol is base 10 (Only valid substrators)
		if i+1 < len(roman) && symbol == 'I' {
			nextSymbol := roman[i+1]

			// buld the two character string
			potentialNumber := string([]byte{symbol, nextSymbol})

			// get the value of the two characater string
			value := allRomanNumerals.ValueOf(potentialNumber)

			if value != 0 {
				total += value
				i++ // move past this character too for next loop
			} else {
				total++
			}
		} else {
			total++
		}
	}
	return total

}
