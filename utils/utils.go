package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ClientCreds struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

type Config struct {
	WaitingTime     int    `yaml:"waiting_time"`
	RedirectionTime int    `yaml:"redirect_time"`
	RedirectUrl     string `yaml:"redirect_url"`
	Jwt             struct {
		Enabled      bool   `yaml:"enabled"`
		PublicKeyUrl string `yaml:"publickey_url"`
	} `yaml:"jwt"`
	Backend struct {
		URL   string `yaml:"server_url"`
		Creds ClientCreds
	} `yaml:"backend"`
}

func (c *Config) AuthURL() string {
	return fmt.Sprintf("%s/auth/v1/", c.Backend.URL)
}

func (c *Config) BackendURL() string {
	return fmt.Sprintf("%s/api/v1/", c.Backend.URL)
}

// LoadConfig Load configuration file
func LoadConfig(configfile string) Config {
	var conf Config
	// Opening config.yml file
	filename, _ := filepath.Abs(configfile)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	// Parsing config file content
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		os.Exit(1)
	}
	return conf
}

// AppendCredentials Append credentials to Config struct
func (c *Config) AppendCreds(credsfils string) {
	creds := make(map[string]interface{})
	filename, _ := filepath.Abs(credsfils)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%s err #%v ", filename, err)
	}
	err = yaml.Unmarshal(file, &creds)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}

	// fmt.Println(creds)
	c.Backend.Creds.ClientID = creds["client_id"].(string)
	c.Backend.Creds.ClientSecret = creds["client_secret"].(string)
}
