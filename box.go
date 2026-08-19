package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(0)

	root := rootCmd()
	root.SetArgs(normaliseArgs(os.Args[1:]))

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func rootCmd() *cobra.Command {
	var noCopy bool

	root := &cobra.Command{
		Use:   "box",
		Short: "Wrap a comment in a box for a given language",
		Long: "Wrap a comment in a box for a given language.\n\n" +
			"The comment is printed to stdout and copied to the clipboard.",
		Example: "  box go This is a comment\n" +
			"  box py This is a comment\n" +
			"  box html This is a comment",
		SilenceUsage: true,
	}

	for _, s := range allStyles {
		root.AddGroup(&cobra.Group{
			ID:    s.id,
			Title: fmt.Sprintf("%s comments:", s.label),
		})
	}

	for _, l := range languages {
		root.AddCommand(languageCmd(l, &noCopy))
	}

	root.AddCommand(languagesCmd())

	root.PersistentFlags().BoolVarP(&noCopy, "no-copy", "n", false, "don't copy the comment to the clipboard")

	return root
}

func languageCmd(l language, noCopy *bool) *cobra.Command {
	return &cobra.Command{
		Use:     l.name + " COMMENT...",
		Aliases: l.aliases,
		Short:   aliasHint(l),
		Long: fmt.Sprintf("Box a %s comment.\n\n%s", l.name,
			l.style.render("This is a comment")),
		GroupID:               l.style.id,
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			box := l.style.render(strings.Join(args, " "))

			fmt.Fprintln(cmd.OutOrStdout(), box)

			if *noCopy {
				return nil
			}

			if err := clipboard.WriteAll(box); err != nil {
				return fmt.Errorf("copying to clipboard: %w", err)
			}

			return nil
		},
	}
}

// aliasHint summarises a language's aliases for the command list, where the
// group heading already makes the comment style obvious.
func aliasHint(l language) string {
	if len(l.aliases) == 0 {
		return ""
	}

	return "aliases: " + strings.Join(l.aliases, ", ")
}

func languagesCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "languages",
		Aliases: []string{"langs"},
		Short:   "List every supported language and its aliases",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "LANGUAGE\tSTYLE\tALIASES")

			for _, l := range languages {
				fmt.Fprintf(w, "%s\t%s\t%s\n", l.name, l.style.label, strings.Join(l.aliases, ", "))
			}

			w.Flush()
		},
	}
}
