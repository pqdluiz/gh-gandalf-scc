package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/pqdluiz/cli-command/analyzer"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gh",
		Short: "GitHub CLI",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("GitHub CLI")
			return nil
		},
	}

	gandalfCmd := &cobra.Command{
		Use:   "scc",
		Short: "Execute a cyclomatic analyzis in your code",
		RunE: func(cmd *cobra.Command, args []string) error {
			directories := []string{"web", "src", "scripts", "internal", "api", "pages", "cmd"}

			err := analyzer.AnalyzeAllFiles(directories)
			if err != nil {
				fmt.Println("Erro:", err)
			}
			return nil
		},
	}

	rootCmd.AddCommand(gandalfCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}

}
