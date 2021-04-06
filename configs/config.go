package configs

import (
	//"fmt"
	"github.com/spf13/viper"
	"os"
)

/* Viper 우선순위
explicit call to Set
flag
env
config
key/value store
default
*/
const (
	envProfileKey     = "profile"
	configLocalName   = "config_local"
	configDevName     = "config_dev"
	configReleaseName = "config_local"
	configFormat      = "yaml"
	configPath        = "resources"
)

func InitConfig() error {

	viper.SetConfigName(getConfigName())
	viper.SetConfigType(configFormat) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv() // AutomaticEnv has Viper check ENV variables for all.
	//println(viper.GetString("DBCONNECTION"))
	return viper.ReadInConfig()
}

func getConfigName() (profileType string) {

	profileType = configLocalName
	if len(os.Getenv(envProfileKey)) > 0 {
		switch key := os.Getenv(envProfileKey); key {
		case "dev":
			profileType = configDevName
			break
		case "release":
			profileType = configReleaseName
		}
	}
	return
}
