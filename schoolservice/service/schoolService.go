package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sathishkumar64/adv_golang/schoolservice/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"

)

func CreateSchool(response http.ResponseWriter, request *http.Request) {
	log.Println("Inside of Create...")
	response.Header().Set("content-type", "application/json")
	var school model.School
	_ = json.NewDecoder(request.Body).Decode(&school)
	result, _ := model.Schooldb.InsertOne(model.MongoCtx, school)
	log.Println("Output of Create...",result)
	json.NewEncoder(response).Encode(result)
}
func GetSchools(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var schools []*model.School

	findOptions := options.Find()
	findOptions.SetLimit(4)
	cursor, err := model.Schooldb.Find(model.MongoCtx, bson.D{{}}, findOptions)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(model.MongoCtx)

	for cursor.Next(model.MongoCtx) {
		var person model.School
		err := cursor.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}
		schools = append(schools, &person)
		//log.Println("The list is....",person)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(schools)

}

func init(){
	model.DbConnect()
}


func main(){
	router := mux.NewRouter()
	//Router handler ,endpoints
	router.HandleFunc("/api/schools", GetSchools).Methods("GET")
	router.HandleFunc("/api/school",CreateSchool).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}