package gotgmd

import (
	"bufio"
	"regexp"
	"strings"
)

func V2(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		line = escape(line)
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func escape(line string) string {
	escapeChars := []string{".", "_", "-", "=", "~", "+", "*", "#", "(", ")", "[", "]", "{", "}", ">", "|", "!", "'", "\""}

	regexes := []*regexp.Regexp{
		regexp.MustCompile(`^\*[^*]+\*$`),          // *bold*
		regexp.MustCompile(`^_[^_]+_$`),            // _italic_
		regexp.MustCompile(`^__[^_]+__$`),          // __underline__
		regexp.MustCompile(`^~[^~]+~$`),            // ~strikethrough~
		regexp.MustCompile(`^\|\|[^|]+\|\|$`),      // ||spoiler||
		regexp.MustCompile(`^\[[^]]+]\([^)]+\)$`),  // [link](url
		regexp.MustCompile(`^!\[[^]]+]\([^)]+\)$`), // ![emoji][link]
	}

	for _, regex := range regexes {
		if regex.MatchString(line) {
			return line
		}
	}

	//Code blocks
	if strings.HasPrefix(line, "```") || strings.Contains(line, "`") {
		return strings.ReplaceAll(line, "\\", "\\\\")
	}

	// Escape the characters that don't match paired constructions or code
	for _, char := range escapeChars {
		line = strings.ReplaceAll(line, char, "\\"+char)
	}

	return line
}
