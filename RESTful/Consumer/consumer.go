package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Course struct {
	Id                        int32  `json:"id"`
	Name                      string `json:"name"`
	Teacher                   int32  `json:"teacher"`
	StudentSatisfactionRating int32  `json:"studentSatisfactionRating,omitempty"`
}

func main() {
	response, err := http.Get("http://localhost:8080/courses")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var courses []Course
	json.Unmarshal([]byte(responseData), &courses)
	for _, element := range courses {
		fmt.Println("Id: ", element.Id)
		fmt.Println("Name: ", element.Name)
		fmt.Println("Teacher: ", element.Teacher)
		fmt.Println("Student Satisfaction Rating: ", element.StudentSatisfactionRating)
		fmt.Println()
	}
}
