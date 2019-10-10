package DBConnector

import (
	"SellerApp/ConfigurationReader"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/golog"
)

var instance *DBConnector

/*GetInstance is a function for creating singleton object of dbconnector*/
func GetInstance() *DBConnector {
	if instance == nil {
		dbConnector := new(DBConnector)
		dbConnector.init()
		instance = dbConnector
	}
	return instance
}

/*DBConnector struct*/
type DBConnector struct {
	Database      *gorm.DB
	configuration *ConfigurationReader.Configuration
}

/*Init function for initializing*/
func (dbConnector *DBConnector) init() {
	dbConnector.configuration = ConfigurationReader.GetInstance()
	dbConnector.openDatabase()
}

func (dbConnector *DBConnector) openDatabase() {

	db, err := gorm.Open("mysql", dbConnector.getConnectionString())
	if err != nil {
		golog.Error("Connection Failed to Open ", err)
		log.Fatal()
	}
	golog.Println("Connection Established")
	dbConnector.Database = db
}

func (dbConnector *DBConnector) getConnectionString() string {
	config := dbConnector.configuration
	return config.DBUserName + `:` + config.DBPassword +
		`@tcp(` + config.DBIP + `:` + config.DBPort + `)/` +
		config.DBName + `?charset=utf8&parseTime=True`
}
