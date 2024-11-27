package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teebow1e/ctfd-useradd/ctfd"
)

var cfgFile string
var csvFilePath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ctfd-useradd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world!")
	},
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Check connection to CTFd as well as the validity of access token.",
	Run: func(cmd *cobra.Command, args []string) {
		ctfd.Ping()
	},
}

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Add user to CTFd using interactive mode.",
	Long:  `In interactive mode, you will be able to create account by entering username, password manually - which gives more control to what account will be created.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Interactive mode started.")
	},
}

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Add user to CTFd with data from a CSV.",
	Long:  `Using this mode, you will be able to add user to CTFd from a CSV, which is often obtained from Google Sheets (as a result of a registration form).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("File mode started.")
		if cmd.Flags().Changed("path") {
			fmt.Printf("Checking file %s\n", csvFilePath)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file (default is config.yaml)")

	fileCmd.Flags().StringVarP(&csvFilePath, "path", "f", "", "Path to CSV file")

	rootCmd.AddCommand(pingCmd)
	rootCmd.AddCommand(interactiveCmd)
	rootCmd.AddCommand(fileCmd)
}
