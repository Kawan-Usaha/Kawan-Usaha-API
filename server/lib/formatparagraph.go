package lib

import "strings"

func FormatParagraph(paragraph string) string {
	// Remove leading and trailing whitespace
	paragraph = strings.TrimSpace(paragraph)

	// Split the paragraph into lines
	lines := strings.Split(paragraph, "\n")

	// Remove empty lines
	var filteredLines []string
	for _, line := range lines {
		if line != "" {
			filteredLines = append(filteredLines, line)
		}
	}

	// Join the filtered lines with newlines
	formattedParagraph := strings.Join(filteredLines, "\n")

	return formattedParagraph
}
