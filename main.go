package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Cars []Car
type Admins []Admin

type Car struct {
    Brand string `json:"Brand"`
    Owner string `json:"Owner"`
    Failure bool `json:"Failure"`
    Repair string `json:"Repair"`
    Price int `json:"Price"`
    arrivalDate string `json:"arrivalDate"`
    departureDate string `json:"departureDate"`
}

type Admin struct {
    Name string `json:"Name"`
    Login string `json:"Login"`
    Password string `json:"Password"`

}

func allAdmins(w http.ResponseWriter, r *http.Request ){
    admins := Admins{
        Admin{Name: "Gautier", Login: "gautierrouxlesang", Password: "wednesdaymydude"},
        Admin{Name: "Helder", Login: "Heldoulabeauté", Password: "wednesdaymydude"},
        Admin{Name: "Vincent", Login: "coucou", Password: "1234"},
    }
    fmt.Println("Endpoint Hit: allAdmins")
    json.NewEncoder(w).Encode(admins)
}

func AllCars(w http.ResponseWriter, r *http.Request){
    cars := Cars{
        Car{Brand:"Toyota", Owner:"M. Dupont", Failure:true, Repair:"Wheels", Price:250},
        Car{Brand:"Peugeot", Owner:"M. Dupond", Failure:true, Repair:"Wheels", Price:250},
        Car{Brand:"Porsche", Owner:"M. Dupompe", Failure:true, Repair:"Wheels", Price:250},
    }
    fmt.Println(r)
    fmt.Println("Endpoint Hit: allCars")
    json.NewEncoder(w).Encode(cars)
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/allcars", AllCars)
    http.HandleFunc("/alladmins", allAdmins)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}