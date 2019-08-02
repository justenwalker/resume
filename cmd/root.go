package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var flags Flags

type Flags struct {
	Filename string
}

var rootCmd = &cobra.Command{
	Use:   "resume",
	Short: "Generates Justen Walker's Resume",
	Long:  `Generates Justen's resume into various formats`,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&flags.Filename, "file", "f", "", "Write the resume out to this file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
