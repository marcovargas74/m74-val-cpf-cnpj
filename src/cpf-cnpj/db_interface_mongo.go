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
		//log.Println(err)
		//w.WriteHeader(http.StatusInternalServerError)
		//fmt.Fprint(w, "Fail to access Mongo DB")
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
		//w.WriteHeader(http.StatusNotFound)
		//fmt.Fprint(w, "DB is Empty")
		errEmpty := errors.New("MONGODB: is Empty")
		return errEmpty.Error(), errEmpty
	}

	json, err := json.Marshal(queryList)
	if err != nil {
		return err.Error(), err
	}

	return string(json), nil

	//w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json")
	//fmt.Fprint(w, string(json))
}

func (q *MyQuery) showQuerysByTypeMongoDB(isCPF bool) (string, error) {

	filter := bson.M{"is_cpf": bson.M{"$eq": isCPF}}

	cursor, err := collectionQuery.Find(ctx, filter)
	if err != nil {
		//log.Println(err)
		//w.WriteHeader(http.StatusInternalServerError)
		//fmt.Fprint(w, "Fail to access Mongo DB")
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
		//w.WriteHeader(http.StatusNotFound)
		//fmt.Fprint(w, "Not Found elements to this Type ")
		errEmpty := errors.New("MONGODB: Not Found elements to this Type")
		return errEmpty.Error(), errEmpty
	}

	/*json, err := json.Marshal(queryList)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}*/

	//w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json")
	//fmt.Fprint(w, string(json))

	json, err := json.Marshal(queryList)
	if err != nil {
		return err.Error(), err
	}

	return string(json), nil

}

/*INCLIR AQUI NOVAS FUNCOES */

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

/*TRASH CODE APAGA DEPOIS */

func (q *MyQuery) showQuerysByNumMongoDB(findNum string) (string, error) {

	filter := bson.M{"cpf": bson.M{"$eq": findNum}}

	cursor, err := collectionQuery.Find(ctx, filter)
	if err != nil {
		//log.Println(err)
		//w.WriteHeader(http.StatusInternalServerError)
		//fmt.Fprint(w, "Fail to access Mongo DB")
		//return
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
		//w.WriteHeader(http.StatusNotFound)
		//fmt.Fprint(w, "Not Found Any elements")
		errEmpty := fmt.Errorf("num %q Not Found", findNum)
		return errEmpty.Error(), errEmpty
	}

	/*json, err := json.Marshal(queryList)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))*/
	json, err := json.Marshal(queryList)
	if err != nil {
		return err.Error(), err
	}

	return string(json), nil
}

/*
func (a *MyQuery) deleteQuerysByNumMongoDB(w http.ResponseWriter, r *http.Request, findNum string) {

	//filter := bson.M{"cpf": bson.M{"$eq": findNum}}

	err := a.deleteQuerysByNumMongoDB_OK(findNum)
	//result, err := collectionQuery.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "SUCCESS TO DELETE CPF/CNPJ")

}

func (a *MyQuery) showQuerysByNumMongoDB(w http.ResponseWriter, r *http.Request, findNum string) {

	filter := bson.M{"cpf": bson.M{"$eq": findNum}}

	cursor, err := collectionQuery.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Fail to access Mongo DB")
		return
	}
	defer cursor.Close(ctx)

	var queryList []MyQuery
	for cursor.Next(ctx) {
		var aQuery MyQuery
		err := cursor.Decode(&aQuery)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		queryList = append(queryList, aQuery)
	}

	if len(queryList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found Any elements")
		return
	}

	json, err := json.Marshal(queryList)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func (a *MyQuery) showQuerysByTypeMongoDB(w http.ResponseWriter, r *http.Request, isCPF bool) {

	filter := bson.M{"is_cpf": bson.M{"$eq": isCPF}}

	cursor, err := collectionQuery.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Fail to access Mongo DB")
		return
	}
	defer cursor.Close(ctx)

	var queryList []MyQuery
	for cursor.Next(ctx) {
		var aQuery MyQuery
		err := cursor.Decode(&aQuery)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		queryList = append(queryList, aQuery)
	}

	if len(queryList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found elements to this Type ")
		return
	}

	json, err := json.Marshal(queryList)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

*/
