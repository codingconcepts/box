package main

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

// kind describes the shape of comment a language uses.
type kind int

const (
	// kindBlock draws a C-style block comment: /* ... */
	kindBlock kind = iota

	// kindLine draws a box out of repeated line comments: # ... #
	kindLine

	// kindWrap draws a box out of comments that open and close on each
	// line: <!-- ... -->
	kindWrap
)

type style struct {
	kind   kind
	id     string
	label  string
	prefix string
	suffix string
	fill   string
}

// lineStyle builds a kindLine style, drawing its borders with the last
// character of the prefix so that they remain valid comments.
func lineStyle(id, prefix string) style {
	return style{
		kind:   kindLine,
		id:     id,
		label:  prefix + " ...",
		prefix: prefix,
		fill:   prefix[len(prefix)-1:],
	}
}

var (
	block = style{kind: kindBlock, id: "block", label: "/* ... */", fill: "*"}
	wrap  = style{kind: kindWrap, id: "wrap", label: "<!-- ... -->", prefix: "<!--", suffix: "-->", fill: "-"}

	hash  = lineStyle("hash", "#")
	slash = lineStyle("slash", "//")
	dash  = lineStyle("dash", "--")
	semi  = lineStyle("semi", ";")
	pct   = lineStyle("pct", "%")
	bang  = lineStyle("bang", "!")
	quote = lineStyle("quote", `"`)
	colon = lineStyle("colon", "::")
)

// allStyles is every supported style, in the order they're listed in help.
var allStyles = []style{block, hash, slash, dash, semi, pct, bang, quote, colon, wrap}

func (s style) render(input string) string {
	switch s.kind {
	case kindWrap:
		return s.renderWrap(input)
	case kindLine:
		return s.renderLine(input)
	default:
		return s.renderBlock(input)
	}
}

func (s style) renderBlock(input string) string {
	border := strings.Repeat(s.fill, utf8.RuneCountInString(input)+3)

	buf := new(bytes.Buffer)
	buf.WriteString("/")
	buf.WriteString(border)
	buf.WriteString("\n")
	buf.WriteString(s.fill)
	buf.WriteString(" ")
	buf.WriteString(input)
	buf.WriteString(" ")
	buf.WriteString(s.fill)
	buf.WriteString("\n")
	buf.WriteString(border)
	buf.WriteString("/")

	return buf.String()
}

func (s style) renderLine(input string) string {
	middle := s.prefix + " " + input + " " + s.prefix
	border := strings.Repeat(s.fill, utf8.RuneCountInString(middle))

	return strings.Join([]string{border, middle, border}, "\n")
}

func (s style) renderWrap(input string) string {
	middle := s.prefix + " " + input + " " + s.suffix
	border := s.prefix + " " + strings.Repeat(s.fill, utf8.RuneCountInString(input)) + " " + s.suffix

	return strings.Join([]string{border, middle, border}, "\n")
}

// normalise makes language lookups forgiving, so that "py", ".py" and ".PY"
// all resolve to the same command.
func normalise(name string) string {
	return strings.ToLower(strings.TrimPrefix(name, "."))
}

// normaliseArgs normalises the language argument, leaving flags and any
// comment text untouched. The first non-flag argument is the language, because
// every flag box defines is a boolean and so never consumes a value.
func normaliseArgs(args []string) []string {
	out := make([]string, len(args))
	copy(out, args)

	for i, arg := range out {
		if strings.HasPrefix(arg, "-") {
			continue
		}

		out[i] = normalise(arg)

		break
	}

	return out
}
