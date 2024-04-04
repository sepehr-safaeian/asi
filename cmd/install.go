package cmd

import (
	"AutoInstaller/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var software config.Software

func installCmd(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("please enter the software name, example: install google-chrome \nif you dont know your software name in AutoInstaller just seen a ./config/softwares.yaml")
		return
	}
	// Open the YAML file
	file, err := os.Open("./config/softwares.yaml")
	if err != nil {
		fmt.Printf("Error opening config file: %v\n", err)
		return
	}
	defer file.Close()

	for _, line := range file {

	}

}
