package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"wyoassign/wyoassign"
	
)


func main() {
	wyoassign.InitAssignments()
	log.Println("starting API server")

	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")


	//specify endpoints
	router.HandleFunc("/api-status", wyoassign.APISTATUS).Methods("GET")
	router.HandleFunc("/assignments", wyoassign.GetAssignments).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.GetAssignment).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.DeleteAssignment).Methods("DELETE")		
	router.HandleFunc("/assignment", wyoassign.CreateAssignment).Methods("POST")	
	router.HandleFunc("/assignments/{id}", wyoassign.UpdateAssignment).Methods("PUT")

	router.HandleFunc("/api-status", wyoassign.APISTATUSS).Methods("GET")
	router.HandleFunc("/cars", wyoassign.GetCars).Methods("GET")
	router.HandleFunc("/car/{type}", wyoassign.GetCar).Methods("GET")
	router.HandleFunc("/car/{type}", wyoassign.DeleteCar).Methods("DELETE")		
	router.HandleFunc("/car", wyoassign.CreateCar).Methods("POST")	
	router.HandleFunc("/carcs/{id}", wyoassign.UpdateCar).Methods("PUT")

	http.Handle("/", router)
	

	//start and listen to requests
	http.ListenAndServe(":8080", router)
	



}