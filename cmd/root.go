// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"net/http"
	"os"

	cobbler "github.com/cobbler/cobblerclient"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var Client cobbler.Client
var conf cobbler.ClientConfig
var httpClient = &http.Client{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobbler",
	Short: "Cobbler CLI client",
	Long:  "An independent CLI to manage a Cobbler server.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobbler.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobbler" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobbler")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// TODO: Do we need the output what configl file is used?
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
	generateCobblerClient()
}

// basic connection to the Cobbler server
func generateCobblerClient() {

	// the configuration is done in .cobbler.yaml
	conf.URL = viper.GetString("server_url")
	conf.Username = viper.GetString("server_username")
	conf.Password = viper.GetString("server_password")

	Client = cobbler.NewClient(httpClient, conf)
	login, _ := Client.Login()

	// TODO: Remove debug messages
	if !login {
		fmt.Println("Login not successful!")
	}
}
