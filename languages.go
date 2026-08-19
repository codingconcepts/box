package main

// language is a subcommand: a canonical name, any number of aliases (usually
// file extensions), and the comment style it renders.
type language struct {
	name    string
	aliases []string
	style   style
}

var languages = []language{
	// Block comments.
	{name: "c", aliases: []string{"h"}, style: block},
	{name: "cpp", aliases: []string{"cc", "cxx", "hpp", "hxx", "cplusplus"}, style: block},
	{name: "csharp", aliases: []string{"cs"}, style: block},
	{name: "css", aliases: []string{"scss", "sass", "less", "styl"}, style: block},
	{name: "d", style: block},
	{name: "dart", style: block},
	{name: "glsl", aliases: []string{"hlsl", "wgsl"}, style: block},
	{name: "go", aliases: []string{"golang"}, style: block},
	{name: "groovy", aliases: []string{"gradle"}, style: block},
	{name: "java", aliases: []string{"pde"}, style: block},
	{name: "javascript", aliases: []string{"js", "cjs", "mjs", "jsx"}, style: block},
	{name: "json5", aliases: []string{"jsonc"}, style: block},
	{name: "kotlin", aliases: []string{"kt", "kts"}, style: block},
	{name: "objectivec", aliases: []string{"objc", "m", "mm"}, style: block},
	{name: "php", style: block},
	{name: "protobuf", aliases: []string{"proto"}, style: block},
	{name: "rust", aliases: []string{"rs"}, style: block},
	{name: "scala", style: block},
	{name: "solidity", aliases: []string{"sol"}, style: block},
	{name: "swift", style: block},
	{name: "typescript", aliases: []string{"ts", "tsx"}, style: block},
	{name: "v", style: block},
	{name: "zig", style: block},

	// Hash comments.
	{name: "awk", style: hash},
	{name: "coffeescript", aliases: []string{"coffee"}, style: hash},
	{name: "conf", aliases: []string{"cfg", "env", "properties"}, style: hash},
	{name: "crystal", aliases: []string{"cr"}, style: hash},
	{name: "docker", aliases: []string{"dockerfile"}, style: hash},
	{name: "elixir", aliases: []string{"ex", "exs"}, style: hash},
	{name: "gdscript", aliases: []string{"gd"}, style: hash},
	{name: "git", aliases: []string{"gitignore", "gitattributes"}, style: hash},
	{name: "julia", aliases: []string{"jl"}, style: hash},
	{name: "make", aliases: []string{"makefile", "mk", "cmake"}, style: hash},
	{name: "nim", style: hash},
	{name: "perl", aliases: []string{"pl", "pm"}, style: hash},
	{name: "powershell", aliases: []string{"ps1", "psm1"}, style: hash},
	{name: "python", aliases: []string{"py", "pyi"}, style: hash},
	{name: "r", style: hash},
	{name: "ruby", aliases: []string{"rb", "rake", "gemfile"}, style: hash},
	{name: "shell", aliases: []string{"sh", "bash", "zsh", "fish", "ksh"}, style: hash},
	{name: "tcl", style: hash},
	{name: "terraform", aliases: []string{"tf", "tfvars", "hcl", "nomad"}, style: hash},
	{name: "toml", style: hash},
	{name: "yaml", aliases: []string{"yml"}, style: hash},

	// Slash comments.
	{name: "jq", style: slash},
	{name: "kdl", style: slash},
	{name: "pkl", style: slash},
	{name: "rego", style: slash},

	// Dash comments.
	{name: "ada", aliases: []string{"adb", "ads"}, style: dash},
	{name: "applescript", aliases: []string{"scpt"}, style: dash},
	{name: "elm", style: dash},
	{name: "haskell", aliases: []string{"hs", "lhs"}, style: dash},
	{name: "lua", style: dash},
	{name: "purescript", aliases: []string{"purs"}, style: dash},
	{name: "sql", aliases: []string{"sqlite"}, style: dash},
	{name: "vhdl", aliases: []string{"vhd"}, style: dash},

	// Semicolon comments.
	{name: "assembly", aliases: []string{"asm", "s"}, style: semi},
	{name: "clojure", aliases: []string{"clj", "cljs", "cljc", "edn"}, style: semi},
	{name: "ini", aliases: []string{"reg", "nsi"}, style: semi},
	{name: "lisp", aliases: []string{"cl", "el", "emacslisp"}, style: semi},
	{name: "racket", aliases: []string{"rkt"}, style: semi},
	{name: "scheme", aliases: []string{"scm", "ss"}, style: semi},

	// Percent comments.
	{name: "erlang", aliases: []string{"erl", "hrl"}, style: pct},
	{name: "latex", aliases: []string{"tex", "sty", "cls", "bib"}, style: pct},
	{name: "matlab", style: pct},
	{name: "prolog", aliases: []string{"pro"}, style: pct},

	// Bang comments.
	{name: "fortran", aliases: []string{"f", "f77", "f90", "f95", "f03", "for"}, style: bang},

	// Quote comments.
	{name: "vim", aliases: []string{"vimrc", "vimscript"}, style: quote},

	// Colon comments.
	{name: "batch", aliases: []string{"bat", "cmd"}, style: colon},

	// Wrapped comments.
	{name: "html", aliases: []string{"htm", "xhtml"}, style: wrap},
	{name: "markdown", aliases: []string{"md"}, style: wrap},
	{name: "svelte", style: wrap},
	{name: "svg", style: wrap},
	{name: "vue", style: wrap},
	{name: "xml", aliases: []string{"xsl", "xslt", "xaml", "plist", "storyboard", "resx", "csproj"}, style: wrap},
}

// languageStyles maps every language name and alias to its style.
var languageStyles = func() map[string]style {
	m := make(map[string]style, len(languages))

	for _, l := range languages {
		m[l.name] = l.style
		for _, alias := range l.aliases {
			m[alias] = l.style
		}
	}

	return m
}()
