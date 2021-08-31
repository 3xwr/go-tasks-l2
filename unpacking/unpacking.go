package unpacking

import (
	"strconv"
	"unicode"
)

type ErrInvalidString struct{}

func (e *ErrInvalidString) Error() string {
	return "invalid string"
}

func UnpackString(input string) (string, error) {

	var escapeSymbol = `\`

	output := ""
	for i := 0; i < len(input); i++ {
		repNum, err := strconv.Atoi(string(input[i]))
		if err != nil {
			if string(input[i]) == escapeSymbol {
				i++
				if i == len(input) {
					break
				}
			}
			output += string(input[i])
			continue
		}
		if i < len(input)-1 && unicode.IsDigit(rune(input[i])) && unicode.IsDigit(rune(input[i+1])) {
			return "", &ErrInvalidString{}
		}
		for j := 0; j < repNum-1; j++ {
			output += string(output[len(output)-1])
		}
	}
	return output, nil
}
