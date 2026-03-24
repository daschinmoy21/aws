package main

import (
	"encoding/json"
	"net/http"
)

type Student struct {
	Name   string `json:"name"`
	RollNo string `json:"roll_number"`
}

func main() {
	http.HandleFunc("/student-details",
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") // For local dev
		json.NewEncoder(w).Encode(Student{
			Name: "Chinmoy", RollNo: "2023bcs52",
		})
	})
	http.ListenAndServe(":8080", nil)
}
