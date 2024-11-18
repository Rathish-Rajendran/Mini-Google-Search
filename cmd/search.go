package cmd

import (
	"fmt"

	"github.com/Rathish-Rajendran/mini-search/search"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches google for the given keyword",
	Run: func(cmd *cobra.Command, args []string) {
		searchHandler(cmd, args)
	},
}

func searchHandler( cmd *cobra.Command, args []string) {
	searchKeyword, err := cmd.Flags().GetString("keyword")
	if err != nil {
		fmt.Errorf("Failed to parse flag %v", err)
	}
	baseSearchUrl := "https://www.google.com/search?q="
	res, err := search.Search(baseSearchUrl, searchKeyword)
	if err != nil {
		fmt.Errorf("Cannot search for the given keyword %v", err)
		return
	}
	if res{
		fmt.Printf("getstream.io is returend when searched for the keyword: %s\n", searchKeyword)
	} else {
		fmt.Printf("getstream.io is not returend when searched for the keyword: %s\n", searchKeyword)
	}
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().String("keyword", "", "The keyword to search")
	searchCmd.MarkFlagRequired("keyword")
}
