package functions

import (
	"fmt"
	"strings"
)

func RenderAscii(input string, asciiMap map[rune][]string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("")
	}

	if len(asciiMap) == 0 {
		return "", fmt.Errorf("font map is empty")
	}

	// Convert \r\n to \n
	input = strings.ReplaceAll(input, "\r\n", "\n")

	var result strings.Builder

	newLine := strings.Split(input, "\n")

	for i, words := range newLine {
		if words == "" {
			if i < len(newLine)-1 {
				result.WriteString("\n")
			}
			continue
		}

		for line := 0; line < 8; line++ {
			for _, ch := range words {
				val, ok := asciiMap[ch]
				if !ok {
					return "", fmt.Errorf("unsupported character: %q", ch)
				}

				result.WriteString(val[line])
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
