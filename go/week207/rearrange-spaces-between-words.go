package week207

import "strings"

func reorderSpaces(text string) string {
	parts := strings.Split(text, " ")
	words := make([]string, 0)
	for _, part := range parts {
		if part != "" {
			words = append(words, part)
		}
	}
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", len(parts)-1)
	}
	join := strings.Repeat(" ", (len(parts)-1)/(len(words)-1))
	return strings.Join(words, join) + strings.Repeat(" ", (len(parts)-1)%(len(words)-1))
}
