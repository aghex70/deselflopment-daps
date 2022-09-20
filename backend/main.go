package main

import (
	"github.com/aghex70/daps/cmd"
	"github.com/aghex70/daps/config"
	"log"
)

func main() {
	log.Println("Starting app...")

	log.Println("Loading configuration...")
	cfg, err := config.NewConfig()
	if err != nil {
		panic("foooooo")
	}
	log.Println("Configuration loaded successfully")
	rootCmd := cmd.RootCommand(cfg)
	err = rootCmd.Execute()
	if err != nil {
		panic("foooooooooooooooo3")
	}

}
