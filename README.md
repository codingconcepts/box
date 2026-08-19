# box

Wraps text from the command line in a box comment for a given language, printing it to stdout and copying it to the clipboard.

### Installation

```sh
go install github.com/codingconcepts/box@latest
```

### Usage

Each language is a subcommand.

```sh
# Use either the languge name.
box LANGUAGE COMMENT...

# Or its extension. 
box EXTENSION COMMENT...
```

Equivalent languages/extension examples.

```sh
$ box python This is a comment
$ box py This is a comment
#####################
# This is a comment #
#####################

$ box javascript This is a comment
$ box js This is a comment
/********************
* This is a comment *
********************/

$ box html This is a comment
<!-- ----------------- -->
<!-- This is a comment -->
<!-- ----------------- -->
```

Langauge names and extensions are case-insensitive and extensions may include a leading dot if that's your thing, so `box .PY ...` works too.

### Commands

```sh
box --help        # every language, grouped by comment style
box python --help # a preview of the box python produces
box languages     # every language, its style, and its aliases
```

### Flags

| Flag | Description |
| ---- | ----------- |
| `-n`, `--no-copy` | Print the comment without copying it to the clipboard |

### Comment styles

| Style | Example languages |
| ----- | ----------------- |
| `/* */` | c, cpp, csharp, css, dart, go, java, javascript, kotlin, php, rust, scala, swift, typescript, zig |
| `#` | conf, docker, elixir, julia, make, perl, powershell, python, ruby, shell, terraform, toml, yaml |
| `--` | ada, applescript, elm, haskell, lua, purescript, sql, vhdl |
| `//` | jq, kdl, pkl, rego |
| `;` | assembly, clojure, ini, lisp, racket, scheme |
| `%` | erlang, latex, matlab, prolog |
| `!` | fortran |
| `"` | vim |
| `::` | batch |
| `<!-- -->` | html, markdown, svelte, svg, vue, xml |

Languages with both line and block comments use the block style, which is why the `//` group holds only languages that have no block comment form.
