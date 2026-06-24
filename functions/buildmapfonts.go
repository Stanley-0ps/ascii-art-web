package functions

import (
	"fmt"
	"os"
	"strings"
)

func BuildFontMap(fileName string) (map[rune][]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("oops, an error has occured while trying to parse file: %w ", err)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("Empty banner file")
	}

	lines := strings.Split(string(data), "\n")
	asciiMaps := make(map[rune][]string)

	if len(lines) != 855 {
		return nil, fmt.Errorf("Incomplete banner file")
	}

	spaceChar := rune(32)

	for i := 0; i+8 <= len(lines); i += 9 {
		asciiMaps[spaceChar] = lines[i+1 : i+9]
		spaceChar++
	}

	if len(asciiMaps) == 0 {
		return nil, fmt.Errorf("oops! failed to build map")
	}
	return asciiMaps, nil
}
