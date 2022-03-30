package wyoassign

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

)

type Response struct{
	Assignments []Assignment `json:"assignments"`
}

type Assignment struct {
	Id string `json:"id"`
	Title string `json:"title`
	Description string `json:"desc"`
	Points int `json:"points"`
}

var Assignments []Assignment
const Valkey string = "FooKey"

type ResponseClass struct{
	Classes []Class `json:"classes"`
}

type Class struct {
	Uid string `json:"uid"`
	Title string `json:"title`
	Teacher string `json:"teacher"`
	NumberOfStudents int `json:"numStudents"`
	NumberOfExams int `json:"numExams"`
}

var Classes []Class

func InitAssignments(){
	var assignmnet Assignment
	assignmnet.Id = "Mike1A"
	assignmnet.Title = "Lab 4 "
	assignmnet.Description = "Some lab this guy made yesteday?"
	assignmnet.Points = 20
	Assignments = append(Assignments, assignmnet)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}


func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assignments = Assignments

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO 
	w.Write(jsonResponse)
}


// func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Entering %s DELETE end point", r.URL.Path)
// 	w.Header().Set("Content-Type", "application/txt")
// 	w.WriteHeader(http.StatusOK)
// 	params := mux.Vars(r)
	
// 	response := make(map[string]string)

// 	response["status"] = "No Such ID to Delete"
// 	for index, assignment := range Assignments {
// 			if assignment.Id == params["id"]{
// 				Assignments = append(Assignments[:index], Assignments[index+1:]...)
// 				response["status"] = "Success"
// 				break
// 			}
// 	}
		
// 	jsonResponse, err := json.Marshal(response)
// 	if err != nil {
// 		return
// 	}
// 	w.Write(jsonResponse)
// }

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	for _, assignment := range Assignments {
		if assignment.Id == params["id"]{
			json.NewEncoder(w).Encode(assignment)
			break
		}
	}
	//TODO : Provide a response if there is no such assignment
	//w.Write(jsonResponse)
}

// <<<<<<< main
// =======
func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, assignment := range Assignments {
			if assignment.Id == params["id"]{
				Assignments = append(Assignments[:index], Assignments[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

// >>>>>>> main
func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	
// <<<<<<< main
	//TODO This should like like cross betweeen Create / Get  
	//w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
//=======
	var response Response
	response.Assignments = Assignments


// >>>>>>> main

	for index, assignment := range Assignments {
		if assignment.Id == params["id"]{
			r.ParseForm()
			if(r.FormValue("id") != ""){
				assignment.Id =  r.FormValue("id")
				assignment.Title =  r.FormValue("title")
				assignment.Description =  r.FormValue("desc")
				assignment.Points, _ =  strconv.Atoi(r.FormValue("points"))
				Assignments = append(Assignments[:index], Assignments[index+1:]...)
				Assignments = append(Assignments, assignment)
				w.WriteHeader(http.StatusOK)
			}
			w.WriteHeader(http.StatusNotFound)
			break
		}
	}
	
	
}

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var assignmnet Assignment
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if(r.FormValue("id") != ""){
		assignmnet.Id =  r.FormValue("id")
		assignmnet.Title =  r.FormValue("title")
		assignmnet.Description =  r.FormValue("desc")
		assignmnet.Points, _ =  strconv.Atoi(r.FormValue("points"))
		Assignments = append(Assignments, assignmnet)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)

}

func InitClasses(){
	var class Class
	class.Uid = "0001"
	class.Title = "Cyber Security"
	class.Teacher = "Roger Rabbit"
	class.NumberOfExams = 2
	class.NumberOfStudents = 20
	Classes = append(Classes, class)
}

func GetClasses(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response ResponseClass

	response.Classes = Classes

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO 
	w.Write(jsonResponse)
}

func GetClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	for _, class := range Classes {
		if class.Uid == params["uid"]{
			json.NewEncoder(w).Encode(class)
			break
		}
	}
	//TODO : Provide a response if there is no such assignment
	//w.Write(jsonResponse)
}

func DeleteClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, class := range Classes {
			if class.Uid == params["uid"]{
				Classes = append(Classes[:index], Classes[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func CreateClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var class Class
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if(r.FormValue("uid") != ""){
		class.Uid =  r.FormValue("uid")
		class.Title =  r.FormValue("title")
		class.Teacher =  r.FormValue("teacher")
		class.NumberOfStudents, _ =  strconv.Atoi(r.FormValue("numStudents"))
		class.NumberOfExams, _ =  strconv.Atoi(r.FormValue("numExams"))
		Classes = append(Classes, class)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)

}