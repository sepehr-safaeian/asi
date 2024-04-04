package cmd

import (
	"AutoInstaller/config"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var softwares []config.Software

func listCmd(cmd *cobra.Command, args []string) {
	filePath := filepath.Join("config", "softwares.yaml")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Softwares.yaml not found:", err)
	}
	defer file.Close()

	// Decoder for YAML
	decoder := yaml.NewDecoder(file)

	// Decode the YAML file into the softwares list
	err = decoder.Decode(&softwares)
	if err != nil {
		log.Fatal("Error decoding softwares.yaml:", err)
	}

	// Print the names of all softwares
	for _, software := range softwares {
		fmt.Println(software.Name)
	}
}
