/*
Copyright © 2023 David Baker

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "GET a random dad joke",
	Long: `This command will GET a random dad joke from icanhazdadjoke.com and print it to your terminal.

Dad jokes are essential to life and children's development.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("random called")
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
