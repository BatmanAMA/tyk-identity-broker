package main

import (
	"encoding/json"
	"github.com/lonelycode/tyk-auth-proxy/tyk-api"
	"io/ioutil"
)

// Configuration holds all configuration settings for TAP
type Configuration struct {
	Secret     string
	Port       int
	ProfileDir string
	BackEnd    struct {
		ProfileBackendSettings  interface{}
		IdentityBackendSettings interface{}
	}
	TykAPISettings tyk.TykAPI
	HttpServerOptions struct {
		UseSSL           bool       
		CertFile string 
		KeyFile  string 
	} 
}

// loadConfig will load the config from a file
func loadConfig(filePath string, configStruct *Configuration) {
	configuration, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error("Couldn't load configuration file: ", err)
		loadConfig("tap.conf", configStruct)
	} else {
		jsErr := json.Unmarshal(configuration, &configStruct)
		if jsErr != nil {
			log.Error("Couldn't unmarshal configuration: ", jsErr)
		}
	}

	log.Debug("[MAIN] Settings Struct: ", configStruct.TykAPISettings)
}
