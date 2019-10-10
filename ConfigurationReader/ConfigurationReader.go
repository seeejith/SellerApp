package ConfigurationReader

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/kataras/golog"
	"github.com/tkanos/gonfig"
)

/*Configuration struct*/
type Configuration struct {
	DBName                string `env:"DBName"`
	DBDriverName          string `env:"DBDriverName"`
	DBUserName            string `env:"DBUserName"`
	DBPassword            string `env:"DBPassword"`
	DBPort                string `env:"DBPort"`
	DBIP                  string `env:"DBIp"`
	MONGODBName           string `env:"MONGODBName"`
	MONGODBDriverName     string `env:"MONGODBDriverName"`
	MONGODBUserName       string `env:"MONGODBUserName"`
	MONGODBPassword       string `env:"MONGODBPassword"`
	MONGODBPort           string `env:"MONGODBPort"`
	MONGODBIP             string `env:"MONGODBIp"`
	MONGODBCOLLECTIONNAME string `env:"MONGODBCollectionName"`
}

var instance *Configuration

/*GetInstance function for retriving the instunce of Configuration*/
func GetInstance() *Configuration {

	if instance == nil {
		instance = readConfiguration()
	}
	return instance
}

func readConfiguration() *Configuration {
	configuration := Configuration{}

	env.Parse(&configuration)
	golog.Println("ENV configuration     ", configuration)

	jsonConfiguration := readConfigurationFromFile()
	instance = &configuration
	if configuration.DBName == "" {
		configuration.DBName = jsonConfiguration.DBName
	}
	if configuration.DBDriverName == "" {
		configuration.DBDriverName = jsonConfiguration.DBDriverName
	}
	if configuration.DBPassword == "" {
		configuration.DBPassword = jsonConfiguration.DBPassword
	}
	if configuration.DBUserName == "" {
		configuration.DBUserName = jsonConfiguration.DBUserName
	}
	if configuration.DBIP == "" {
		configuration.DBIP = jsonConfiguration.DBIP
	}
	if configuration.DBPort == "" {
		configuration.DBPort = jsonConfiguration.DBPort
	}
	if configuration.MONGODBName == "" {
		configuration.MONGODBName = jsonConfiguration.MONGODBName
	}
	if configuration.MONGODBDriverName == "" {
		configuration.MONGODBDriverName = jsonConfiguration.MONGODBDriverName
	}
	if configuration.MONGODBPassword == "" {
		configuration.MONGODBPassword = jsonConfiguration.MONGODBPassword
	}
	if configuration.MONGODBUserName == "" {
		configuration.MONGODBUserName = jsonConfiguration.MONGODBUserName
	}
	if configuration.MONGODBIP == "" {
		configuration.MONGODBIP = jsonConfiguration.MONGODBIP
	}
	if configuration.MONGODBPort == "" {
		configuration.MONGODBPort = jsonConfiguration.MONGODBPort
	}
	return &configuration
}

func readConfigurationFromEnv() Configuration {
	configuration := Configuration{}
	err := env.Parse(&configuration)
	if err != nil {
		golog.Error("reading env veriable failed", err)
	} else {
		golog.Println("ENV configuration     ", configuration)
	}
	return configuration
}

func readConfigurationFromFile() Configuration {
	jsonConfiguration := Configuration{}

	err := gonfig.GetConf("configuration.json", &jsonConfiguration)
	if err != nil {
		golog.Error("reading configuration failed ", err)
		log.Fatal()
	}
	golog.Println("jsonConfiguration    ", jsonConfiguration)
	return jsonConfiguration
}
