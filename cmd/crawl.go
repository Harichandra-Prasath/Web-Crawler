package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var crawlcmd = &cobra.Command{
	Use:   "crawl",
	Short: "Main subcommand of crawl..Entrypoint",
	RunE: func(cmd *cobra.Command, args []string) error {
		root_url, _ := cmd.Flags().GetString("url")
		depth, _ := cmd.Flags().GetInt("depth")
		same_domain, _ := cmd.Flags().GetBool("same-domain")
		to_generate, _ := cmd.Flags().GetBool("generate")
		if depth == -1 {
			return fmt.Errorf("depth is required for crawling")
		}
		if root_url == "" {
			return fmt.Errorf("root url is needed")
		} else {
			_, err := http.Get(root_url)
			if err != nil {
				fmt.Printf("Error in root url")
				return err
			}
			Crawl(root_url, same_domain, to_generate, depth)

		}
		return nil
	},
}

func init() {
	rootcmd.AddCommand(crawlcmd)
	crawlcmd.PersistentFlags().String("url", "", "Root url to start the scraping")
	crawlcmd.PersistentFlags().Bool("same-domain", false, "Crawling and scraping from same domains")
	crawlcmd.PersistentFlags().Bool("generate", false, "Generate a .txt file with all the links crawled")
	crawlcmd.PersistentFlags().Int("depth", -1, "Defines the depth level of crawling including root")

}
