package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "Start The AutoInstaller",
	Long:  "a fast and secure auto installer software & codec on windows (currently).",
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "help",
		Short: "Display help information",
		Long:  `Display help information for the application. This command will display a list of all available commands and their descriptions.`,
		Run:   helpCmd,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "list all already software to install",
		Run:   listCmd,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "install",
		Short: "install software",
		Run:   installCmd,
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
