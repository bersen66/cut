package cmd

import (
	"log"

	"github.com/bersen66/cut/pkg/filter"
	"github.com/spf13/cobra"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// used in flags
var (
	delimeter string
	separator string
	field     int
	rootCmd   = &cobra.Command{
		Use:   "cut [flags]",
		Short: "Cut out fields from stdin or file",
		Long:  "Floppa - big russian cat",
		Run: func(cmd *cobra.Command, args []string) {
			config := &filter.Config{
				Separator: separator,
				Delimeter: delimeter,
				Field:     field,
			}
			err := filter.Run(config)
			checkError(err)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&field, "fields", "f", 0, "select cut field to cut")
	rootCmd.PersistentFlags().StringVarP(&delimeter, "delimeter", "d", "\t", "specify delimeter between columns")
	rootCmd.PersistentFlags().StringVarP(&separator, "separator", "s", NewLineString, "specify separator between lines")

	err := rootCmd.MarkPersistentFlagRequired("fields")
	checkError(err)
}
