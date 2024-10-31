// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"encoding/json"
	"fmt"
	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"log/slog"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

var cfgFile string
var Client cobbler.Client
var conf cobbler.ClientConfig
var httpClient = &http.Client{}
var verbose bool

// NewRootCmd builds a new command that represents the base action when called without any subcommands
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "cobbler",
		Short: "Cobbler CLI client",
		Long:  "An independent CLI to manage a Cobbler server.",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	// global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobbler.yaml)")
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Whether or not to print debug messages from the CLI.")

	// Add sub commands
	rootCmd.AddCommand(NewAclSetupCmd())
	rootCmd.AddCommand(NewBuildisoCmd())
	distroCmd, err := NewDistroCmd()
	cobra.CheckErr(err)
	rootCmd.AddCommand(distroCmd)
	rootCmd.AddCommand(NewEventCmd())
	rootCmd.AddCommand(NewFileCmd())
	rootCmd.AddCommand(NewHardlinkCmd())
	rootCmd.AddCommand(NewImageCmd())
	rootCmd.AddCommand(NewImportCmd())
	rootCmd.AddCommand(NewListCmd())
	rootCmd.AddCommand(NewMenuCmd())
	rootCmd.AddCommand(NewMgmtClassCmd())
	rootCmd.AddCommand(NewMkLoadersCmd())
	rootCmd.AddCommand(NewPackageCmd())
	rootCmd.AddCommand(NewProfileCmd())
	rootCmd.AddCommand(NewReplicateCmd())
	rootCmd.AddCommand(NewRepoCmd())
	rootCmd.AddCommand(NewReportCmd())
	rootCmd.AddCommand(NewRepoSyncCmd())
	rootCmd.AddCommand(NewSettingCmd())
	rootCmd.AddCommand(NewSignatureCmd())
	rootCmd.AddCommand(NewSyncCmd())
	rootCmd.AddCommand(NewSystemCmd())
	rootCmd.AddCommand(NewValidateAutoinstallsCmd())
	rootCmd.AddCommand(NewVersionCmd())
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()

	// Execute root command
	cobra.CheckErr(rootCmd.Execute())
}

func setupLogger() {
	if !verbose {
		slog.SetLogLoggerLevel(slog.LevelWarn)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Set defaults
	viper.SetDefault("server_url", "http://127.0.0.1/cobbler_api")
	viper.SetDefault("server_username", "cobbler")
	viper.SetDefault("server_password", "cobbler")

	// Read config file
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
	err := viper.ReadInConfig()
	if cfgFile != "" {
		cobra.CheckErr(err)
	}
	if verbose {
		_, _ = fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
	}
}

// basic connection to the Cobbler server
func generateCobblerClient() error {

	// the configuration is done in .cobbler.yaml
	conf.URL = viper.GetString("server_url")
	conf.Username = viper.GetString("server_username")
	conf.Password = viper.GetString("server_password")

	Client = cobbler.NewClient(httpClient, conf)
	login, err := Client.Login()

	if !login {
		return fmt.Errorf("failed to login")
	}
	return err
}

func printStructured(cmd *cobra.Command, dataStruct interface{}) {
	s := reflect.ValueOf(dataStruct).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		mapstructureTag := typeOfT.Field(i).Tag.Get("mapstructure")
		fieldName := typeOfT.Field(i).Name
		fieldStructName := typeOfT.Field(i).Type.String()
		if strings.HasPrefix(fieldStructName, "cobblerclient.Value") {
			printValueStructured(cmd, mapstructureTag, f)
			continue
		}
		if fieldName == "Item" {
			baseItem := f.Interface().(cobbler.Item)
			printStructured(cmd, &baseItem)
			continue
		}
		if fieldName == "Interfaces" {
			// Skip and print at the end
			continue
		}
		if fieldName == "Client" {
			continue
		}
		if fieldName == "Meta" {
			continue
		}
		printField(cmd, f.Kind(), mapstructureTag, f.Interface())
	}

	// Print interfaces at the end of the output
	networkInterfacesField := s.FieldByName("Interfaces")
	if networkInterfacesField != (reflect.Value{}) {
		networkInterfaces := networkInterfacesField.Interface().(cobbler.Interfaces)
		printNetworkInterface(cmd, networkInterfaces)
	}
}

func printValueStructured(cmd *cobra.Command, name string, value reflect.Value) {
	isInherited := value.FieldByName("IsInherited").Bool()
	data := value.FieldByName("Data").Interface()
	if isInherited {
		printField(cmd, reflect.String, name, "<<inherit>>")
	} else {
		dataType := value.FieldByName("Data").Kind()
		printField(cmd, dataType, name, data)
	}
}

func printNetworkInterface(cmd *cobra.Command, networkInterface cobbler.Interfaces) {
	for interfaceName, interfaceStruct := range networkInterface {
		fmt.Fprintf(cmd.OutOrStdout(), "%-40s: %s\n", "Interface =====", interfaceName)
		printStructured(cmd, &interfaceStruct)
	}
}

func printField(cmd *cobra.Command, valueType reflect.Kind, name string, value interface{}) {
	if name == "ctime" || name == "mtime" {
		time, err := covertFloatToUtcTime(value.(float64))
		if err == nil {
			// If there is an error just show the float
			fmt.Fprintf(cmd.OutOrStdout(), "%-40s: %s\n", name, time)
			return
		}
	}
	switch valueType {
	case reflect.Bool:
		fmt.Fprintf(cmd.OutOrStdout(), "%-40s: %t\n", name, value.(bool))
	case reflect.Int64:
		fmt.Fprintf(cmd.OutOrStdout(), "%-40s: %d\n", name, value.(int64))
	case reflect.Int32:
		fmt.Fprintf(cmd.OutOrStdout(), "%-40s: %d\n", name, value.(int32))
	case reflect.Int16:
		fmt.Fprintf(cmd.OutOrStdout(), "%-40s: %d\n", name, value.(int16))
	case reflect.Int8:
		fmt.Fprintf(cmd.OutOrStdout(), "%-40s: %d\n", name, value.(int8))
	case reflect.Int:
		fmt.Fprintf(cmd.OutOrStdout(), "%-40s: %d\n", name, value.(int))
	case reflect.Float32:
		fmt.Fprintf(cmd.OutOrStdout(), "%-40s: %f\n", name, value.(float32))
	case reflect.Float64:
		fmt.Fprintf(cmd.OutOrStdout(), "%-40s: %f\n", name, value.(float64))
	case reflect.Map:
		res2B, _ := json.Marshal(value)
		fmt.Fprintf(cmd.OutOrStdout(), "%-40s: %s\n", name, string(res2B))
	case reflect.Array, reflect.Slice:
		arr := reflect.ValueOf(value)
		fmt.Fprintf(cmd.OutOrStdout(), "%-40s: [", name)
		for i := 0; i < arr.Len(); i++ {
			if i+1 != arr.Len() {
				fmt.Fprintf(cmd.OutOrStdout(), "'%v', ", arr.Index(i).Interface())
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), "'%v'", arr.Index(i).Interface())
			}
		}
		fmt.Fprintf(cmd.OutOrStdout(), "]\n")
	default:
		if value == nil {
			value = ""
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%-40s: %s\n", name, value)
		// fmt.Fprintf(cmd.OutOrStdout(),"%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
