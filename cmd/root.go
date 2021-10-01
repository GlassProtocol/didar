package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "didar",
	Short: "A mvp did spec built on arweave",
	Long:  `do this later`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.didar.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().String("arweave-key", "", "path to arweave key file")
	rootCmd.MarkPersistentFlagRequired("arweave-key")
	viper.BindPFlag("arweave-key", rootCmd.PersistentFlags().Lookup("arweave-key"))

	rootCmd.PersistentFlags().String("protocol", "", "which protocol are we using")
	rootCmd.MarkPersistentFlagRequired("protocol")
	viper.BindPFlag("protocol", rootCmd.PersistentFlags().Lookup("protocol"))

	rootCmd.PersistentFlags().String("public-key", "", "pub key")
	rootCmd.MarkPersistentFlagRequired("public-key")
	viper.BindPFlag("public-key", rootCmd.PersistentFlags().Lookup("public-key"))

	rootCmd.PersistentFlags().String("private-key", "", "raw text for private key (dangerous change me)")
	rootCmd.MarkPersistentFlagRequired("private-key")
	viper.BindPFlag("private-key", rootCmd.PersistentFlags().Lookup("private-key"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".didar" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".didar")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
