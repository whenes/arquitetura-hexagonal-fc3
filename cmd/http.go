package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/codeedu/go-hexagonal/adapters/cli"
	server2 "github.com/codeedu/go-hexagonal/adapters/web/server"
)

var action string
var productId string
var productName string
var productPrice float64

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your command. For example:

Cobra is a CLI application framework for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := server2.MakeNewWebServer()
		server.Service = &application.ProductService
		fmt.Println("Webserver has been started")
		server.Serve()
}

func init() {
	rootCmd.AddCommand(httpCmd)

	httpCmd.Flags().StringVarP(&action, "action", "a", "enable", "Enable/Disable a product")
	httpCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")
	httpCmd.Flags().StringVarP(&productName, "product", "n", "", "Product Name")
	httpCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "Product Price")
}