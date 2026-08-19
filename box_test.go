package main

import (
	"reflect"
	"strings"
	"testing"
)

type testCase struct {
	name  string
	lang  string
	input string
	exp   string
}

// styleCases contains one case per supported comment style. Every style must
// appear here; TestEveryStyleHasCase enforces it.
func styleCases() []testCase {
	return []testCase{
		{
			name:  "block style",
			lang:  "go",
			input: "This is a comment",
			exp: "/********************\n" +
				"* This is a comment *\n" +
				"********************/",
		},
		{
			name:  "hash style",
			lang:  "py",
			input: "This is a comment",
			exp: "#####################\n" +
				"# This is a comment #\n" +
				"#####################",
		},
		{
			name:  "slash style",
			lang:  "rego",
			input: "This is a comment",
			exp: "///////////////////////\n" +
				"// This is a comment //\n" +
				"///////////////////////",
		},
		{
			name:  "dash style",
			lang:  "sql",
			input: "This is a comment",
			exp: "-----------------------\n" +
				"-- This is a comment --\n" +
				"-----------------------",
		},
		{
			name:  "semicolon style",
			lang:  "clj",
			input: "This is a comment",
			exp: ";;;;;;;;;;;;;;;;;;;;;\n" +
				"; This is a comment ;\n" +
				";;;;;;;;;;;;;;;;;;;;;",
		},
		{
			name:  "percent style",
			lang:  "tex",
			input: "This is a comment",
			exp: "%%%%%%%%%%%%%%%%%%%%%\n" +
				"% This is a comment %\n" +
				"%%%%%%%%%%%%%%%%%%%%%",
		},
		{
			name:  "bang style",
			lang:  "f90",
			input: "This is a comment",
			exp: "!!!!!!!!!!!!!!!!!!!!!\n" +
				"! This is a comment !\n" +
				"!!!!!!!!!!!!!!!!!!!!!",
		},
		{
			name:  "quote style",
			lang:  "vim",
			input: "This is a comment",
			exp: `"""""""""""""""""""""` + "\n" +
				`" This is a comment "` + "\n" +
				`"""""""""""""""""""""`,
		},
		{
			name:  "colon style",
			lang:  "bat",
			input: "This is a comment",
			exp: ":::::::::::::::::::::::\n" +
				":: This is a comment ::\n" +
				":::::::::::::::::::::::",
		},
		{
			name:  "wrap style",
			lang:  "html",
			input: "This is a comment",
			exp: "<!-- ----------------- -->\n" +
				"<!-- This is a comment -->\n" +
				"<!-- ----------------- -->",
		},
		{
			name:  "multi-byte input is measured in runes",
			lang:  "go",
			input: "héllo",
			exp: "/********\n" +
				"* héllo *\n" +
				"********/",
		},
	}
}

func TestRender(t *testing.T) {
	for _, c := range styleCases() {
		t.Run(c.name, func(t *testing.T) {
			equals(t, c.exp, styleFor(t, c.lang).render(c.input))
		})
	}
}

// TestEveryStyleHasCase fails if a style is added without a corresponding case
// in styleCases.
func TestEveryStyleHasCase(t *testing.T) {
	covered := map[string]bool{}
	for _, c := range styleCases() {
		covered[styleFor(t, c.lang).id] = true
	}

	for _, s := range allStyles {
		if !covered[s.id] {
			t.Fatalf("no case in styleCases covers the %q style", s.id)
		}
	}
}

// TestRenderWidths asserts that every supported language produces a three-line
// box whose lines are all the same width.
func TestRenderWidths(t *testing.T) {
	for name, s := range languageStyles {
		t.Run(name, func(t *testing.T) {
			act := s.render("This is a comment")

			lines := strings.Split(act, "\n")
			if len(lines) != 3 {
				t.Fatalf("\n\texp: 3 lines\n\tgot: %d", len(lines))
			}

			for _, line := range lines[1:] {
				if len(line) != len(lines[0]) {
					t.Fatalf("\n\tinconsistent line widths in:\n%s", act)
				}
			}
		})
	}
}

func styleFor(tb testing.TB, lang string) style {
	tb.Helper()

	s, ok := languageStyles[lang]
	if !ok {
		tb.Fatalf("unregistered language %q", lang)
	}

	return s
}

func equals(tb testing.TB, exp any, act any) {
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("\n\texp: %#[1]v (%[1]T)\n\tgot: %#[2]v (%[2]T)", exp, act)
	}
}
