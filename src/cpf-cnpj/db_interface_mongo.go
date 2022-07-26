package cpfcnpj

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	//DBMongo_Local Const used to Open Local db
	DBMongo_Local  = "mongodb://localhost:27017"
	DBMongo_Docker = "mongodb://root:example@mongo:27017/"
	SetDockerRun   = true
	SetLocalRun    = false
)

//IsUsingMongoDocker If Using MongoDB in  a Docker image this var is True
var IsUsingMongoDocker bool

var collectionQuery *mongo.Collection
var ctx = context.TODO()

//GetIsUsingMongoDocker Get If Using MongoDB in  a Docker image
func GetIsUsingMongoDocker() bool {
	return IsUsingMongoDocker
}

//SetUsingMongoDocker set If Using MongoDB in  a Docker image
func SetUsingMongoDocker(isMongoDocker bool) {
	IsUsingMongoDocker = isMongoDocker
}

//InitDBMongo Initi MOngo Database
func InitDBMongo(isDockerRun bool) bool {
	urlMongo := DBMongo_Local
	if isDockerRun {
		urlMongo = DBMongo_Docker
	}
	clientOptions := options.Client().ApplyURI(urlMongo)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return false
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("connecting to the MONGO DB...")
	collectionQuery = client.Database("Querys").Collection("Querys")
	fmt.Println("Successfully connected to the DB MONGO!")
	return true

}

func (q *MyQuery) saveQueryInMongoDB() error {

	_, err := collectionQuery.InsertOne(ctx, q)
	return err
}

func (q *MyQuery) showQueryAllMongoDB() (string, error) {

	filter := bson.M{}

	cursor, err := collectionQuery.Find(ctx, filter)
	if err != nil {
		return err.Error(), err
	}
	defer cursor.Close(ctx)

	var queryList []MyQuery
	for cursor.Next(ctx) {
		var aQuery MyQuery
		err := cursor.Decode(&aQuery)
		if err != nil {
			return err.Error(), err
		}
		queryList = append(queryList, aQuery)
	}

	if len(queryList) == 0 {
		errEmpty := errors.New("MONGODB: is Empty")
		return errEmpty.Error(), errEmpty
	}

	json, err := json.Marshal(queryList)
	if err != nil {
		return err.Error(), err
	}

	return string(json), nil
}

func (q *MyQuery) showQuerysByTypeMongoDB(isCPF bool) (string, error) {

	filter := bson.M{"is_cpf": bson.M{"$eq": isCPF}}

	cursor, err := collectionQuery.Find(ctx, filter)
	if err != nil {
		return err.Error(), err
	}
	defer cursor.Close(ctx)

	var queryList []MyQuery
	for cursor.Next(ctx) {
		var aQuery MyQuery
		err := cursor.Decode(&aQuery)
		if err != nil {
			return err.Error(), err
		}
		queryList = append(queryList, aQuery)
	}

	if len(queryList) == 0 {
		errEmpty := errors.New("MONGODB: Not Found elements to this Type")
		return errEmpty.Error(), errEmpty
	}

	json, err := json.Marshal(queryList)
	if err != nil {
		return err.Error(), err
	}

	return string(json), nil

}

func (q *MyQuery) deleteQuerysByNumMongoDB(findNum string) error {

	filter := bson.M{"cpf": bson.M{"$eq": findNum}}

	result, err := collectionQuery.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("num %q Not Found", findNum)
	}

	return nil
}

func (q *MyQuery) showQuerysByNumMongoDB(findNum string) (string, error) {

	filter := bson.M{"cpf": bson.M{"$eq": findNum}}

	cursor, err := collectionQuery.Find(ctx, filter)
	if err != nil {
		return err.Error(), err
	}
	defer cursor.Close(ctx)

	var queryList []MyQuery
	for cursor.Next(ctx) {
		var aQuery MyQuery
		err := cursor.Decode(&aQuery)
		if err != nil {
			return err.Error(), err
		}
		queryList = append(queryList, aQuery)
	}

	if len(queryList) == 0 {
		errEmpty := fmt.Errorf("num %q Not Found", findNum)
		return errEmpty.Error(), errEmpty
	}

	json, err := json.Marshal(queryList)
	if err != nil {
		return err.Error(), err
	}

	return string(json), nil
}
