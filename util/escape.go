package util

import (
	"fmt"
	"strings"
)

func Escape(s string, parseMode string) string {
	switch parseMode {
	case "HTML":
		return EscapeHtml(s)
	case "Markdown":
		return EscapeMarkdown(s)
	case "MarkdownV2":
		return EscapeMarkdownV2(s)
	default:
		panic(fmt.Sprintf("Escape: unsupported mode %q, supported modes are: HTML, Markdown, MarkdownV2", parseMode))
	}
}

var htmlReplacer = strings.NewReplacer(
	"<", "&lt;",
	">", "&gt;",
	"&", "&amp;",
	"\"", "&quot;",
)

func EscapeHtml(s string) string {
	return htmlReplacer.Replace(s)
}

var markdownReplacer = strings.NewReplacer(
	"_", "\\_",
	"*", "\\*",
	"`", "\\`",
	"[", "\\[",
)

func EscapeMarkdown(s string) string {
	return markdownReplacer.Replace(s)
}

var markdownV2Replacer = strings.NewReplacer(
	"_", "\\_",
	"*", "\\*",
	"[", "\\[",
	"]", "\\]",
	"(", "\\(",
	")", "\\)",
	"~", "\\~",
	"`", "\\`",
	">", "\\>",
	"#", "\\#",
	"+", "\\+",
	"-", "\\-",
	"=", "\\=",
	"|", "\\|",
	"{", "\\{",
	"}", "\\}",
	".", "\\.",
	"!", "\\!",
	"\\", "\\\\",
)

func EscapeMarkdownV2(s string) string {
	return markdownV2Replacer.Replace(s)
}
