package main

import (
	"context"
	"log"
	"net"

	"mAssignment1/course"

	"google.golang.org/grpc"
)

type server struct {
	course.UnimplementedCoursesServer
}

//the type of data we can handle
type Course struct {
	Id                        int32  `json:"id"`
	Name                      string `json:"name"`
	Teacher                   int32  `json:"teacher"`
	StudentSatisfactionRating int32  `json:"studentSatisfactionRating,omitempty"`
}

//dummy data
var courses = []Course{
	{Id: 1, Name: "BDSA", Teacher: 111, StudentSatisfactionRating: 5},
	{Id: 2, Name: "IDBS", Teacher: 222, StudentSatisfactionRating: 2},
	{Id: 3, Name: "DSYS", Teacher: 333, StudentSatisfactionRating: 0},
}

func (s *server) GetCourseList(ctx context.Context, in *course.GetCourseListRequest) (*course.GetCourseListReply, error) {
	//implement actual logic of getting the courses
	var coursesReply = []*course.Course{
		{Id: 1, Name: "BDSA", Teacher: 111, StudentSatisfactionRating: 5},
		{Id: 2, Name: "IDBS", Teacher: 222, StudentSatisfactionRating: 2},
		{Id: 3, Name: "DSYS", Teacher: 333, StudentSatisfactionRating: 0},
	}
	return &course.GetCourseListReply{Message: "Here goes the list of courses ;)", Courses: coursesReply}, nil //.HelloReply{Message: fmt.Sprint(courses[num].Name) + ";" + fmt.Sprint(courses[num].StudentSatisfactionRating)}, nil
}

func main() {
	//port :8080
	lis, err := net.Listen("tcp", ":8080")

	if err != nil { //error before listening
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //we create a new server
	course.RegisterCoursesServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil { //error while listening
		log.Fatalf("failed to serve: %v", err)
	}
}
