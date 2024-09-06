package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var (
	default_lbx          string
	default_db           string
	default_yaml_db_path string
	yaml_db              string
)

func load_config() {

	// Load configuration from config.yml file
	// viper.SetConfigName("config") // name of config file (without extension)
	// viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name

	loadDefaults()

	viper.SetDefault("yaml_db", default_yaml_db_path+strings.Replace(filepath.Base(os.Args[0]), ".exe", "", -1)+".yml-db")

}

func loadDefaults() {
	viper.SetDefault("default_lbx", "default_lbx not set")
	viper.SetDefault("default_db", "default_db not set")
	viper.SetDefault("default_yaml_db_path", "default_yaml_db_path not set")

	// Load configuration from config.yml file
	// viper.SetConfigName("config") // name of config file (without extension)
	// viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

			fmt.Println("📢 Config file not found! Using and writing defaults to 'config.yml' ...")
			viper.WriteConfigAs("config.yml")
		} else {
			panic(fmt.Errorf("📢 fatal error config file: %w", err))
		}
	}

	default_lbx = viper.GetString("default_lbx")
	default_db = viper.GetString("default_db")
	default_yaml_db_path = viper.GetString("default_yaml_db_path")
}
