package asciiart

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func GenerateAsciiArt(input, font string) (string, error) {
	filePath := "fonts/" + font + ".txt"
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("%s banner does not exist", font)
	}

	content := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	output := &strings.Builder{}
	PrintAsciiArt(strings.Split(input, "\n"), lines, output)
	return output.String(), nil
}

func PrintAsciiArt(sentences []string, textFile []string, output io.Writer) {
	for _, word := range sentences {
		if word == "" {
			fmt.Fprintln(output)
			continue
		}
		for h := 1; h < 9; h++ {
			for i := 0; i < len(word); i++ {
				for lineIndex, line := range textFile {
					if lineIndex == (int(word[i])-32)*9+h {
						fmt.Fprint(output, line)
					}
				}
			}
			fmt.Fprintln(output)
		}
	}
}
