package main

import (
	"context"
	"log"

	"mAssignment1/course"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock()) //maybe it has to be: localhost:8080
	if err != nil {                                                        //error can not establish connection
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := course.NewCoursesClient(conn)

	response, err := client.GetCourseList(context.Background(), &course.GetCourseListRequest{})
	if err != nil {
		log.Fatalf("could not get course list: %v", err)
	}

	log.Printf("The response from server with list of courses is: %s", response.GetMessage())
	log.Println(response.GetCourses())
}
