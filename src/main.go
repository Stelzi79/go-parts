package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Hello World!")

	fmt.Println("🪄 Loading configuration ...")
	load_config()

	for key, value := range viper.AllSettings() {
		fmt.Printf("%s='%v'\n", key, value)

	}
	// fmt.Println(viper.AllSettings())
}
