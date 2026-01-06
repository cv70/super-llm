package llm

import "strings"

func ParseOutput(s string) (string, string) {
	thinkEndIdx := strings.Index(s, "</think>")
	if thinkEndIdx < 0 {
		return "", s
	}
	thinkEndIdx += len("</think>")
	return s[:thinkEndIdx], s[thinkEndIdx:]
}

func RemoveThink(s string) string {
	_, s = ParseOutput(s)
	return s
}
