package cpfcnpj

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	//DBMongo_Local Const used to Open Local db
	DBMongo_Local  = "mongodb://localhost:27017"
	DBMongo_Docker = "mongodb://root:example@mongo:27017/"
)

var collectionQuery *mongo.Collection
var ctx = context.TODO()

//InitDBMongo Initi MOngo Database
func InitDBMongo(isDockerRun bool) bool {
	urlMongo := DBMongo_Local
	if isDockerRun {
		//urlMongo = os.Getenv("MONGO_URL")
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

	//defer db.Close()

	fmt.Println("connecting to the MONGO DB...")
	collectionQuery = client.Database("Querys").Collection("Querys")
	fmt.Println("Successfully connected to the DB MONGO!")
	return true

}

func (q *MyQuery) saveQueryInMongoDB() bool {

	_, err := collectionQuery.InsertOne(ctx, q)
	if err != nil {
		log.Println(err)
		return true
	}
	return true
}

func (q *MyQuery) showQueryAllMongoDB(w http.ResponseWriter, r *http.Request) {

	filter := bson.M{}

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
		fmt.Fprint(w, "DB is Empty")
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

func (a *MyQuery) deleteQuerysByNumMongoDB(w http.ResponseWriter, r *http.Request, findNum string) {

	filter := bson.M{"cpf": bson.M{"$eq": findNum}}

	result, err := collectionQuery.DeleteOne(ctx, filter)
	//result, err := collectionQuery.DeleteMany(ctx, filter)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Fail to access DB")
		return
	}

	if result.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found elements to this Type - Delete")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "SUCCESS TO DELETE CPF/CNPJ")

}
