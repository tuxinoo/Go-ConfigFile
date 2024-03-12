package main

import (
	"fmt"
	"os"
	helper "sample/utils"

	"gopkg.in/yaml.v2"
)

const (
	ConfigurationFile string = "config.yaml"
	CredentialsFile   string = "creds.yaml"
)

var config helper.Config

func init() {
	config = helper.LoadConfig(ConfigurationFile)
	config.AppendCreds(CredentialsFile)
}

func debug() {
	// Dump YAML formatted file
	dump, err := yaml.Marshal(&config)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("**** Configuration values ****\n%s", string(dump))
}

func main() {
	// debug()
	fmt.Println(config.AuthURL())
	fmt.Println(config.BackendURL())
}
