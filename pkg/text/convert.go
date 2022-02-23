package text

import "unicode"

// Copied from:
// https://github.com/asaskevich/govalidator/blob/f21760c49a8d602d863493de796926d2a5c1138d/utils.go#L107
func CamelToKebab(str string) string {
	var output []rune
	var segment []rune
	for _, r := range str {
		// not treat number as separate segment
		if !unicode.IsLower(r) && string(r) != "-" && !unicode.IsNumber(r) {
			output = addSegment(output, segment)
			segment = nil
		}
		segment = append(segment, unicode.ToLower(r))
	}
	output = addSegment(output, segment)
	return string(output)
}

func addSegment(inrune, segment []rune) []rune {
	if len(segment) == 0 {
		return inrune
	}
	if len(inrune) != 0 {
		inrune = append(inrune, '-')
	}
	inrune = append(inrune, segment...)
	return inrune
}
