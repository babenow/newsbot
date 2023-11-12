package markup

import "strings"

var (
	replacer = strings.NewReplacer(
		"-",
		"\\-",
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
		"=",
		"\\=",
	)
)

func EscapeForMarkdown(str string) string {
	return replacer.Replace(str)
}
