package main

import (
	"fmt"

	"github.com/DisguisedTrolley/gator/internal/config"
)

func main() {
	cfgFile, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	cfgFile.SetUser("Samarth")

	cfgFile, err = config.ReadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Read config: %+v\n", cfgFile)
}
