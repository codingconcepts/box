package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRootPrintsHelpWithoutArgs(t *testing.T) {
	out, err := execute(t)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, exp := range []string{"Usage:", "box [command]", "languages"} {
		if !strings.Contains(out, exp) {
			t.Fatalf("\n\texp help containing: %q\n\tgot:\n%s", exp, out)
		}
	}
}

func TestLanguageCommands(t *testing.T) {
	cases := []struct {
		name string
		args []string
		exp  string
	}{
		{
			name: "canonical name",
			args: []string{"python", "This", "is", "a", "comment"},
			exp: "#####################\n" +
				"# This is a comment #\n" +
				"#####################\n",
		},
		{
			name: "alias",
			args: []string{"py", "This", "is", "a", "comment"},
			exp: "#####################\n" +
				"# This is a comment #\n" +
				"#####################\n",
		},
		{
			name: "words are joined into a single comment",
			args: []string{"go", "This", "is", "a", "comment"},
			exp: "/********************\n" +
				"* This is a comment *\n" +
				"********************/\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out, err := execute(t, append([]string{"--no-copy"}, c.args...)...)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			equals(t, c.exp, out)
		})
	}
}

func TestLanguageCommandRequiresComment(t *testing.T) {
	if _, err := execute(t, "--no-copy", "go"); err == nil {
		t.Fatal("\n\texp error for missing comment\n\tgot: nil")
	}
}

func TestUnknownLanguage(t *testing.T) {
	_, err := execute(t, "nope", "This is a comment")
	if err == nil {
		t.Fatal("\n\texp error for unknown language\n\tgot: nil")
	}

	if !strings.Contains(err.Error(), "nope") {
		t.Fatalf("\n\texp error mentioning %q\n\tgot: %q", "nope", err.Error())
	}
}

// TestEveryLanguageHasCommand asserts the language table and the command tree
// stay in step, and that no name or alias is registered twice.
func TestEveryLanguageHasCommand(t *testing.T) {
	root := rootCmd()

	seen := map[string]string{}

	for _, l := range languages {
		for _, name := range append([]string{l.name}, l.aliases...) {
			if owner, dupe := seen[name]; dupe {
				t.Fatalf("%q is registered by both %q and %q", name, owner, l.name)
			}

			seen[name] = l.name

			if name != normalise(name) {
				t.Fatalf("%q is not normalised, so it can never be matched", name)
			}

			cmd, _, err := root.Find([]string{name})
			if err != nil {
				t.Fatalf("finding command %q: %v", name, err)
			}

			if cmd.Name() != l.name {
				t.Fatalf("\n\texp %q to resolve to %q\n\tgot: %q", name, l.name, cmd.Name())
			}
		}
	}
}

// TestEveryStyleHasLanguage asserts no style is defined without a language
// using it, which would leave an empty group in the help output.
func TestEveryStyleHasLanguage(t *testing.T) {
	used := map[string]bool{}
	for _, l := range languages {
		used[l.style.id] = true
	}

	for _, s := range allStyles {
		if !used[s.id] {
			t.Fatalf("no language uses the %q style", s.id)
		}
	}
}

func TestNormaliseArgs(t *testing.T) {
	cases := []struct {
		name string
		args []string
		exp  []string
	}{
		{
			name: "no args",
			args: []string{},
			exp:  []string{},
		},
		{
			name: "language is normalised",
			args: []string{".GO", "This", "is", "a", "comment"},
			exp:  []string{"go", "This", "is", "a", "comment"},
		},
		{
			name: "language after a flag is normalised",
			args: []string{"-n", ".PY", "This", "is", "a", "comment"},
			exp:  []string{"-n", "py", "This", "is", "a", "comment"},
		},
		{
			name: "comment text is left alone",
			args: []string{"go", "This.Is", "A", "COMMENT"},
			exp:  []string{"go", "This.Is", "A", "COMMENT"},
		},
		{
			name: "flags only",
			args: []string{"--help"},
			exp:  []string{"--help"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			equals(t, c.exp, normaliseArgs(c.args))
		})
	}
}

// execute runs the root command with args, returning anything it wrote.
func execute(tb testing.TB, args ...string) (string, error) {
	tb.Helper()

	buf := new(bytes.Buffer)

	root := rootCmd()
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	err := root.Execute()

	return buf.String(), err
}
