package options

import (
	"encoding/json"
	"os"
	
	"k8s.io/klog"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

func LoadConfig() *Configuration {
	file, err := os.Open("chapter02-chitchat/options/config.json")
	if err != nil {
		klog.Fatal("open config file error: %v", err)
	}
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		klog.Fatal("decode from file error: %v", err)
	}
	return &config
}

// Version
func Version() string {
	return "0.1"
}
