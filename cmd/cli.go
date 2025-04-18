package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/codeedu/go-hexagonal/adapters/cli"
)

var action string
var productId string
var productName string
var productPrice float64

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your command. For example:

Cobra is a CLI application framework for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(&productService, action, productId, productName, productPrice)
		if err != nil {
			fmt.Println("Error:", err.Error())
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "Enable/Disable a product")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")
	cliCmd.Flags().StringVarP(&productName, "product", "n", "", "Product Name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "Product Price")
}