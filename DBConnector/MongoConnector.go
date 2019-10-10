package DBConnector

import (
	"SellerApp/ConfigurationReader"
	"fmt"
	"log"
	"github.com/kataras/golog"
	"github.com/go-bongo/bongo"
)

type MongoConnector struct {
	Database      *bongo.Connection
	configuration *ConfigurationReader.Configuration
}

var mongoInstance *MongoConnector

func GetMongoInstance() *MongoConnector {
	if mongoInstance == nil {
		mongoConnector := new(MongoConnector)
		mongoConnector.init()
		mongoInstance = mongoConnector
	}
	return mongoInstance
}

func (connector *MongoConnector) init() {
	connector.configuration = ConfigurationReader.GetInstance()
	connector.Database = connector.openDarabase()
}

func (connector *MongoConnector) getConfig() *bongo.Config {
	config := &bongo.Config{
		ConnectionString: "localhost",
		Database:         "bongotest1",
	}
	return config
}

func (connector *MongoConnector) openDarabase() *bongo.Connection {
	config := connector.getConfig()
	connection, err := bongo.Connect(config)
	if err != nil {
		golog.Error("Error while create connection to MONGODB ", err)
		log.Fatal(err)
	}
	err = connection.Connect()
	if err != nil {
		golog.Error("Error while connecting MONGODB ", err)
		log.Fatal(err)
	}
	fmt.Println(connection)
	return connection
}

func (connector *MongoConnector) createConnectionString() string {
	config := connector.configuration
	connectionString := config.MONGODBDriverName + `://` + config.MONGODBIP + `:` +
		config.MONGODBPort

	fmt.Println(connectionString)
	return connectionString
}
