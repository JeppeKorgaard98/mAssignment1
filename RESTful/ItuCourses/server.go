package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Course struct {
	Id                        int32  `json:"id"`
	Name                      string `json:"name"`
	Teacher                   int32  `json:"teacher"`
	StudentSatisfactionRating int32  `json:"studentSatisfactionRating,omitempty"`
}

var courses = []Course{
	{Id: 1, Name: "BDSA", Teacher: 111, StudentSatisfactionRating: 5},
	{Id: 2, Name: "IDBS", Teacher: 222, StudentSatisfactionRating: 2},
	{Id: 3, Name: "DSYS", Teacher: 333, StudentSatisfactionRating: 0},
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)

	router.Run("localhost:8080")

}
