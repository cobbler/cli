// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"encoding/json"
	"fmt"
	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobbler",
	Short: "Cobbler CLI client",
	Long:  "An independent CLI to manage a Cobbler server.",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
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
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Whether or not to print debug messages from the CLI.")
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
	if err := viper.ReadInConfig(); err == nil {
		// TODO: Do we need the output what config file is used?
		if verbose {
			_, _ = fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
		}
	}
}

// basic connection to the Cobbler server
func generateCobblerClient() {

	// the configuration is done in .cobbler.yaml
	conf.URL = viper.GetString("server_url")
	conf.Username = viper.GetString("server_username")
	conf.Password = viper.GetString("server_password")

	Client = cobbler.NewClient(httpClient, conf)
	login, err := Client.Login()

	if !login || err != nil {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf("error! Failed to login: %s", err))
	}
}

// simply prints a message about functions not implemented in the cobblerclient library
func notImplemented() {
	_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf(`error! Not yet implemented in the cobblerclient library
See https://github.com/cobbler/cobblerclient/issues/4`))
}

func printStructured(dataStruct interface{}) {
	s := reflect.ValueOf(dataStruct).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		mapstructureTag := typeOfT.Field(i).Tag.Get("mapstructure")
		fieldName := typeOfT.Field(i).Name
		fieldStructName := typeOfT.Field(i).Type.String()
		if strings.HasPrefix(fieldStructName, "cobblerclient.Value") {
			printValueStructured(mapstructureTag, f)
			continue
		}
		if fieldName == "Item" {
			baseItem := f.Interface().(cobbler.Item)
			printStructured(&baseItem)
			continue
		}
		if fieldName == "Interfaces" {
			baseItem := f.Interface().(map[string]interface{})
			printNetworkInterface(baseItem)
			continue
		}
		if fieldName == "Client" {
			continue
		}
		if fieldName == "Meta" {
			continue
		}
		printField(f.Kind(), mapstructureTag, f.Interface())
	}
}

func printValueStructured(name string, value reflect.Value) {
	isInherited := value.FieldByName("IsInherited").Bool()
	data := value.FieldByName("Data").Interface()
	if isInherited {
		printField(reflect.String, name, "<<inherit>>")
	} else {
		dataType := value.FieldByName("Data").Kind()
		printField(dataType, name, data)
	}
}

func printInterfaceStructured(data interface{}) {
	v, ok := data.(map[string]interface{})
	if !ok {
		panic("Cast of interface to map unsuccessful during structured fmt.Printf!")
	}
	for key, value := range v {
		printField(reflect.ValueOf(value).Kind(), key, value)
	}
}

func printNetworkInterface(networkInterface map[string]interface{}) {
	for interfaceName, interfaceStruct := range networkInterface {
		fmt.Printf("%-40s: %s\n", "Interface =====", interfaceName)
		printInterfaceStructured(interfaceStruct)
	}
}

func printField(valueType reflect.Kind, name string, value interface{}) {
	if name == "ctime" || name == "mtime" {
		time, err := covertFloatToUtcTime(value.(float64))
		if err == nil {
			// If there is an error just show the float
			fmt.Printf("%-40s: %s\n", name, time)
			return
		}
	}
	switch valueType {
	case reflect.Bool:
		fmt.Printf("%-40s: %t\n", name, value.(bool))
	case reflect.Int64:
		fmt.Printf("%-40s: %d\n", name, value.(int64))
	case reflect.Int32:
		fmt.Printf("%-40s: %d\n", name, value.(int32))
	case reflect.Int16:
		fmt.Printf("%-40s: %d\n", name, value.(int16))
	case reflect.Int8:
		fmt.Printf("%-40s: %d\n", name, value.(int8))
	case reflect.Int:
		fmt.Printf("%-40s: %d\n", name, value.(int))
	case reflect.Float32:
		fmt.Printf("%-40s: %f\n", name, value.(float32))
	case reflect.Float64:
		fmt.Printf("%-40s: %f\n", name, value.(float64))
	case reflect.Map:
		res2B, _ := json.Marshal(value)
		fmt.Printf("%-40s: %s\n", name, string(res2B))
	case reflect.Array, reflect.Slice:
		arr := reflect.ValueOf(value)
		fmt.Printf("%-40s: [", name)
		for i := 0; i < arr.Len(); i++ {
			if i+1 != arr.Len() {
				fmt.Printf("'%v', ", arr.Index(i).Interface())
			} else {
				fmt.Printf("'%v'", arr.Index(i).Interface())
			}
		}
		fmt.Printf("]\n")
	default:
		if value == nil {
			value = ""
		}
		fmt.Printf("%-40s: %s\n", name, value)
		// fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
