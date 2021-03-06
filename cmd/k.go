package cmd

import (
	"fmt"
	"gogit/pkg/base"
	"gogit/pkg/data"

	"github.com/spf13/cobra"
)

var kCmd = &cobra.Command{
	Use:   "k",
	Short: "Prints Refs",
	Run: func(cmd *cobra.Command, args []string) {
		dot := "digraph commits {\n"

		var oids []string
		for refname, ref := range data.IterRefs(false) {
			dot += fmt.Sprintf("\"%s\" [shape=note]\n", refname)
			dot += fmt.Sprintf("\"%s\" -> \"%s\"\n", refname, ref)
			if ref.Symbolic == false {
				oids = append(oids, ref.Value)
			}
		}

		for _, oid := range base.IterCommitsAndParents(oids) {
			commit := base.GetCommit(oid)
			dot += fmt.Sprintf("\"%s\" [shape=box style=filled label=\"%s\"]\n", oid, oid[:10])
			if commit.Parent != "" {
				dot += fmt.Sprintf("\"%s\" -> \"%s\"\n", oid, commit.Parent)
			}
		}

		dot += "}"

		print(dot)
	},
}
