package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func helpCmd(cmd *cobra.Command, args []string) {
	fmt.Println("**Help Information**")
	fmt.Println()
	fmt.Println("This application provides various tools for managing software.")
	fmt.Println("Use the following commands to explore its functionalities:")
	fmt.Println()
	fmt.Println("* `help` - Displays detailed information about all available commands.")
	fmt.Println("* `search <software>` - Searches for software within the authorized installer list based on the provided keyword.")
	fmt.Println("* `install <software>` - Installs the specified software from the allowed installer list.")
	fmt.Println()
	fmt.Println("For specific help on a command, use `help <command name>`.")
}
