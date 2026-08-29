package util

import (
	"strings"
	"testing"
)

func TestEscapeHTML(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"plain text", "hello world", "hello world"},
		{"ampersand", "Tom & Jerry", "Tom &amp; Jerry"},
		{"less than", "1 < 2", "1 &lt; 2"},
		{"greater than", "2 > 1", "2 &gt; 1"},
		{"all specials", "<b>a & b</b>", "&lt;b&gt;a &amp; b&lt;/b&gt;"},
		{"quotes", `he said "hi"`, `he said &quot;hi&quot;`},
		{"apostrophe untouched", "it's fine", "it's fine"},
		{"already escaped entity", "&amp;", "&amp;amp;"},
		{"repeated specials", "<<>>&&", "&lt;&lt;&gt;&gt;&amp;&amp;"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := EscapeHtml(tc.in)
			if got != tc.want {
				t.Errorf("EscapeHTML(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestEscapeMarkdown(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"plain text", "hello world", "hello world"},
		{"underscore", "snake_case", `snake\_case`},
		{"asterisk", "bold*text*", `bold\*text\*`},
		{"backtick", "code`snippet`", "code\\`snippet\\`"},
		{"open bracket", "link[text]", `link\[text]`},
		{"all specials", "_*`[", "\\_\\*\\`\\["},
		{"close bracket untouched", "text]", "text]"},
		{"multiple underscores", "__bold__", `\_\_bold\_\_`},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := EscapeMarkdown(tc.in)
			if got != tc.want {
				t.Errorf("EscapeMarkdown(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestEscapeMarkdownV2(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"plain text", "hello world", "hello world"},
		{"underscore", "a_b", `a\_b`},
		{"asterisk", "a*b", `a\*b`},
		{"brackets", "a[b]c", `a\[b\]c`},
		{"parens", "a(b)c", `a\(b\)c`},
		{"tilde", "a~b", `a\~b`},
		{"backtick", "a`b", "a\\`b"},
		{"gt", "a>b", `a\>b`},
		{"hash", "a#b", `a\#b`},
		{"plus", "a+b", `a\+b`},
		{"minus", "a-b", `a\-b`},
		{"equals", "a=b", `a\=b`},
		{"pipe", "a|b", `a\|b`},
		{"braces", "a{b}c", `a\{b\}c`},
		{"dot", "a.b", `a\.b`},
		{"exclamation", "a!b", `a\!b`},
		{"backslash", `a\b`, `a\\b`},
		{"all specials together", "_*[]()~`>#+-=|{}.!",
			"\\_\\*\\[\\]\\(\\)\\~\\`\\>\\#\\+\\-\\=\\|\\{\\}\\.\\!"},
		{"unicode text", "Привет, мир!", `Привет, мир\!`},
		{"mixed text and specials", "Price: $5.00 (was $10.00)!",
			`Price: $5\.00 \(was $10\.00\)\!`},
		{"emoji untouched", "hello 😀 world", "hello 😀 world"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := EscapeMarkdownV2(tc.in)
			if got != tc.want {
				t.Errorf("EscapeMarkdownV2(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

// Benchmarks to make sure escaping stays cheap on longer inputs.
func BenchmarkEscapeHTML(b *testing.B) {
	s := strings.Repeat("<tag> & \"text\" ", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EscapeHtml(s)
	}
}

func BenchmarkEscapeMarkdown(b *testing.B) {
	s := strings.Repeat("_*`[text] ", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EscapeMarkdown(s)
	}
}

func BenchmarkEscapeMarkdownV2(b *testing.B) {
	s := strings.Repeat("_*[]()~`>#+-=|{}.! text ", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EscapeMarkdownV2(s)
	}
}
