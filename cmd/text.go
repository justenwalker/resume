package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/justenwalker/resume/data"

	"github.com/Masterminds/sprig"
	"github.com/spf13/cobra"
)

var textTemplate string

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Ouput Resume in Text format",
	RunE: func(cmd *cobra.Command, args []string) error {
		tmpl, err := template.New(filepath.Base(textTemplate)).Funcs(template.FuncMap(sprig.FuncMap())).ParseFiles(textTemplate)
		if err != nil {
			return fmt.Errorf("can't parse text template file %s: %v", textTemplate, err)
		}

		fd, err := os.Create(flags.Filename)
		if err != nil {
			return fmt.Errorf("can't create file %s: %v", flags.Filename, err)
		}
		defer fd.Close()
		if err := tmpl.Execute(fd, data.MyCV); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(textCmd)
	textCmd.PersistentFlags().StringVarP(&textTemplate, "template", "t", "", "Use the given go text/template file")
}
