package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	homedir "github.com/mitchellh/go-homedir"
	dbInfra "github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/spf13/viper"
)

var cfgFile string

var db, _ = sql.Open("sqlite3", "db.sqlite")
var productDb = dbInfra.NewProductDb(db)
var productService = application.ProductService{Persistence: productDb}

var rootCmd = &cobra.Command{
	Use:   "go-hexagonal",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI application framework for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName(".go-hexagonal")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
