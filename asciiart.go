package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var outputFileWritten bool // Global flag to track if output has been written

// Automatically save the ASCII art output to a file named "output.txt"
func saveToOutput(asciiChars [][]string, asciiHeight int) {
	outputFile := "output.txt"

	// Step 1: If this is the first time writing in a new generation, open in create/overwrite mode
	var file *os.File
	var err error
	if !outputFileWritten {
		file, err = os.Create(outputFile) // Overwrite the file for a new ASCII generation
		if err != nil {
			fmt.Printf("Error: Could not create or open the file: %v\n", err)
			return
		}
		outputFileWritten = true // Mark that the file has been written to
	} else {
		file, err = os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0644) // Append for subsequent lines
		if err != nil {
			fmt.Printf("Error: Could not open the file for appending: %v\n", err)
			return
		}
	}
	defer file.Close()

	// Step 2: Convert the asciiChars (2D array) to a string
	var asciiArtContent strings.Builder
	for i := 0; i < asciiHeight; i++ {
		for j := 0; j < len(asciiChars); j++ {
			asciiArtContent.WriteString(asciiChars[j][i]) // Add each line of the character
			if j < len(asciiChars)-1 {
				asciiArtContent.WriteString(" ") // Optional space between characters
			}
		}
		asciiArtContent.WriteString("\n") // Move to the next line of the ASCII art
	}

	// Step 3: Write the content to the file
	_, err = file.WriteString(asciiArtContent.String())
	if err != nil {
		fmt.Printf("Error: Failed to write to the file: %v\n", err)
		return
	}
}

// ProcessString processes the input string, returns ASCII art, and saves it to output.txt
func ProcessString(input string, asciiMap map[rune][]string, asciiHeight int) string {
	log.Printf("Processing input string: %s", input)
	input = strings.ReplaceAll(input, `\n`, "\n")
	inputLines := strings.Split(input, "\n")
	var result string

	outputFileWritten = false // Reset flag to overwrite output file for a new generation

	// Process each line separately
	for _, line := range inputLines {
		if line == "" {
			result += "\n" // Handle empty lines (newline in the ASCII art)
			continue
		}
		// Build and append the ASCII art for the line
		asciiChars := buildAsciiArt(line, asciiMap, asciiHeight)
		asciiArtLine := generateAsciiArt(asciiChars, asciiHeight)
		result += asciiArtLine
		saveToOutput(asciiChars, asciiHeight) // Save each processed line to output.txt
	}

	return result
}

// Build the ASCII art for a given line of input
func buildAsciiArt(line string, asciiMap map[rune][]string, asciiHeight int) [][]string {
	var asciiChars [][]string
	for _, char := range line { // Use rune to support Unicode
		if art, exists := asciiMap[char]; exists {
			asciiChars = append(asciiChars, art)
		} else {
			asciiChars = append(asciiChars, make([]string, asciiHeight))
		}
	}
	return asciiChars
}

// Generate the ASCII art string for the given characters
func generateAsciiArt(asciiChars [][]string, asciiHeight int) string {
	var result string

	for i := 0; i < asciiHeight; i++ {
		for _, charLines := range asciiChars {
			result += charLines[i] // Append each line of the ASCII character
		}
		result += "\n" // Move to the next line of the ASCII art
	}

	return result
}

// Load the ASCII art banner into a map
func LoadBanner(filename string) (map[rune][]string, int, error) {
	log.Printf("Loading banner from file: %s", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Failed to load banner: %v", err)
		return nil, 0, fmt.Errorf("failed to load font: %v", err)
	}
	strData := strings.ReplaceAll(string(data), "\r\n", "\n")

	// Check if the first character in the data is a newline and skip it if necessary
	if len(strData) > 0 && strData[0] == '\n' {
		strData = strData[1:]
	}

	asciiLines := strings.Split(strData, "\n\n")
	asciiMap := make(map[rune][]string)

	// Create a map for ASCII characters (starting from the space character)
	for i, art := range asciiLines {
		char := rune(32 + i) // ASCII code starts at 32 for space
		asciiMap[char] = strings.Split(art, "\n")
	}

	// Determine asciiHeight dynamically
	var asciiHeight int
	for _, art := range asciiMap {
		asciiHeight = len(art)
		break // Get the height from the first character
	}
	log.Printf("Loaded banner with height: %d", asciiHeight)
	return asciiMap, asciiHeight, nil
}

// Check the validity of the input (ensure it's ASCII)
func CheckValidity(input string) bool {
	for _, r := range input {
		if r > unicode.MaxASCII {
			log.Printf("Invalid character detected: %c", r)
			return false
		}
	}
	return true
}
