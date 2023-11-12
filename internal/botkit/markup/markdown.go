package markup

import "strings"

var (
	replacer = strings.NewReplacer(
		"-",
		"\\-",
		"_",
		"*",
		"\\*",
		"[",
		"\\[",
		"]",
		"\\]",
		"(",
		"\\(",
		")",
		"\\)",
		"{",
		"\\{",
		"}",
		"\\}",
		".",
		"\\.",
		"!",
		"\\!",
	)
)

func EscapeForMarkdown(str string) string {
	return replacer.Replace(str)
}
