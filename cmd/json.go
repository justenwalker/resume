package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/justenwalker/resume/data"

	"github.com/spf13/cobra"
)

var jsonIndent bool

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "Ouput Resume in json format",
	RunE: func(cmd *cobra.Command, args []string) error {
		fd, err := os.Create(flags.Filename)
		if err != nil {
			return fmt.Errorf("can't create file %s: %v", flags.Filename, err)
		}
		defer fd.Close()
		enc := json.NewEncoder(fd)
		enc.SetEscapeHTML(false)
		enc.SetIndent("", "  ")
		return enc.Encode(data.MyCV)
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)
	jsonCmd.PersistentFlags().BoolVar(&jsonIndent, "format", false, "Indent output json for readability.")
}
