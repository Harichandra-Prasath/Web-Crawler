package cmd

// subcommand "crawl"

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var crawlcmd = &cobra.Command{
	Use:   "crawl",
	Short: "Main subcommand of crawler..Entrypoint",
	RunE: func(cmd *cobra.Command, args []string) error {
		root_url, _ := cmd.Flags().GetString("url")
		depth, _ := cmd.Flags().GetInt("depth")
		root_relative, _ := cmd.Flags().GetBool("root-relative")
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
			// As there is no usage of url librarys for validation
			// we have to mainpulate strings to get the best results

			if root_url[len(root_url)-1] != '/' { // This is additonal logic to get relative paths
				root_url += "/"
			}
			Intiate(root_url, root_relative, to_generate, depth)

		}
		return nil
	},
}

func init() {
	rootcmd.AddCommand(crawlcmd)
	crawlcmd.PersistentFlags().String("url", "", "Root url to start the scraping")
	crawlcmd.PersistentFlags().Bool("root-relative", false, "Crawling and scraping from same domain with root path")
	crawlcmd.PersistentFlags().Bool("generate", false, "Generate a .txt file with all the links crawled")
	crawlcmd.PersistentFlags().Int("depth", -1, "Defines the depth level of crawling including root")

}
