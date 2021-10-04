package main

import (
	"context"
	"log"
	"net"

	"mAssignment1/course" ///Users/Freja/Documents/Programmering/3.semester/disys/Mandatory1/mAssignment1

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

func (s *server) GetCourseList(ctx context.Context, in *course.GetCourseListRequest) (*course.GetCourseListReply, error) {
	//implement actual logic of getting the courses
	return &course.GetCourseListReply{Message: "Here goes the list of courses ;)"}, nil //.HelloReply{Message: fmt.Sprint(courses[num].Name) + ";" + fmt.Sprint(courses[num].StudentSatisfactionRating)}, nil
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
