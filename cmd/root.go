package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "CryptoPulse",
	Short: "CryptoPulse is a CLI tool for getting cryptocurrency prices",
	Long: `CryptoPulse is a CLI tool for getting cryptocurrency prices.
	You can use it to get the latest prices of cryptocurrencies.`,
	Run: func(cmd *cobra.Command, args []string) {
		// show help short message
		cmd.Help()
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
