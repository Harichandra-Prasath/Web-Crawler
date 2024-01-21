package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootcmd = &cobra.Command{
	Use:   "crawler",
	Short: "Crawl all the links",
	RunE: func(cmd *cobra.Command, args []string) error {
		if ok, _ := cmd.Flags().GetBool("version"); ok {
			fmt.Println("version 1.0")
			return nil
		}
		return nil
	},
}

func Run() error {
	err := rootcmd.Execute()
	return err
}

func init() {
	rootcmd.PersistentFlags().Bool("version", false, "Flag to chek the version of the crawler")
}
