package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func load_config() {

	viper.SetDefault("default_lbx", "default_lbx not set")
	viper.SetDefault("default_db", "default_db not set")

	// Load configuration from config.yml file
	// viper.SetConfigName("config") // name of config file (without extension)
	// viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name

	viper.AddConfigPath("../") // optionally look for config in the working directory
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("📢 Config file not found! Using and writing defaults to 'config.yml' ...")
			viper.WriteConfigAs("config.yml")
		} else {
			panic(fmt.Errorf("📢 fatal error config file: %w", err)) // Config file was found but another error was produced
		}
	}

}
