package UserHandler

import (
	"SellerApp/DBConnector"
	"encoding/json"
	"fmt"
	"time"

	"github.com/fatih/structs"
	"github.com/go-bongo/bongo"
	"github.com/kataras/golog"
)


type UserHandler struct {
	dbConnecor *DBConnector.DBConnector
}

type Person struct {
	bongo.DocumentBase `bson:",inline"`
	Data               map[string]interface{}
}

var userHandlerInstance *UserHandler

func GetUserHandlerInstance() *UserHandler {
	if userHandlerInstance == nil {
		userHandler := new(UserHandler)
		userHandler.init()
		userHandlerInstance = userHandler
	}
	return userHandlerInstance
}

func (userHandler *UserHandler) init() {
	golog.Println("init ")
	userHandler.dbConnecor = DBConnector.GetInstance()
	userHandler.dbConnecor.Database.Debug().AutoMigrate(&Profile{})
	userHandler.dbConnecor.Database.Debug().AutoMigrate(&Auth{})
	userHandler.dbConnecor.Database.Debug().AutoMigrate(&Access{})
	userHandler.insertDummyData()
}

func (userHandler *UserHandler) insertDummyData() {

	db := userHandler.dbConnecor.Database

	profile := new(Profile)
	profile.Name = "sreejithh"
	profile.Email = "sreee@gmail.com"
	profile.Age = 27
	profile.MobileNumber = "9898989898"
	db.Create(profile)

	access := new(Access)
	access.ProfileId = profile.Id
	db.Create(access)

	auth := new(Auth)
	auth.ProfileId = profile.Id
	auth.LastLoginTime = time.Now()
	db.Create(auth)

}

func (userHandler *UserHandler) TransformUserData(profile Profile) {
	db := userHandler.dbConnecor.Database
	profiles := make([]Profile, 0)
	// profile := new(Profile)
	// profile.Name = "sreejith"
	db.Where(&profile).Find(&profiles)
	for _, p := range profiles {
		profileMap := userHandler.getProfileData(p)
		err := userHandler.writeDataTOmongo(profileMap)
		if err != nil {
			golog.Println("failed to write to mongo ", profileMap)
		}
	}
}

func (userHandler *UserHandler) createQuery(profile Profile) map[string]interface{} {
	profileMap := make(map[string]interface{}, 0)
	byteData, _ := json.Marshal(profile)
	fmt.Println("byteData ", string(byteData))
	json.Unmarshal(byteData, &profileMap)
	query := make(map[string]interface{}, 0)
	for k, v := range profileMap {
		if k == "Id" {
			continue
		}
		query[`data.`+k] = v
	}
	return query
}
func (userHandler *UserHandler) getProfileData(profile Profile) map[string]interface{} {

	profileMap := structs.Map(profile)
	authMap := userHandler.getAuth(profile)
	dataMap := mergeMap(profileMap, authMap)
	accessMap := userHandler.getAccess(profile)
	dataMap = mergeMap(dataMap, accessMap)
	return dataMap
}

func (userHandler *UserHandler) getAuth(profile Profile) map[string]interface{} {
	db := userHandler.dbConnecor.Database
	auth := new(Auth)
	auth.ProfileId = profile.Id
	db.Where(&auth).First(&auth)
	return structs.Map(auth)
}

func (userHandler *UserHandler) getAccess(profile Profile) map[string]interface{} {
	db := userHandler.dbConnecor.Database
	access := new(Access)
	access.ProfileId = profile.Id
	db.Where(&access).First(&access)
	return structs.Map(access)
}

func mergeMap(map1 map[string]interface{}, map2 map[string]interface{}) map[string]interface{} {
	for k, v := range map2 {
		map1[k] = v
	}
	return map1
}

func (userHandler *UserHandler) GetUserData(profile Profile) (map[string]interface{}, error) {
	mongo := DBConnector.GetMongoInstance()
	person := new(Person)
	query := userHandler.createQuery(profile)
	fmt.Println("s  ", query)
	err := mongo.Database.Collection("test").FindOne(query, &person)
	if err != nil {
		golog.Println("data reading from mogo failed ", err)
		return nil, err
	}

	golog.Println("testPerson ", person)
	return person.Data, nil
}

func (userHandler *UserHandler) writeDataTOmongo(data map[string]interface{}) error {
	person := new(Person)
	person.Data = data
	mongo := DBConnector.GetMongoInstance()
	err := mongo.Database.Collection("test").Save(person)
	if err != nil {
		golog.Println("err ", err)
	}
	return err
}
