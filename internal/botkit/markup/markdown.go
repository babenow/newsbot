package markup

import "strings"

var (
	replacer = strings.NewReplacer(
		"-",
		"\\-",
		"_",
		"\\_",
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
		"~",
		"\\~",
		"`",
		"\\`",
		">",
		"\\>",
		"#",
		"\\#",
		"+",
		"\\+",
		"=",
		"\\=",
		"|",
		"\\|",
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
