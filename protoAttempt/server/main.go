package protoAttempt //package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	pb "https://github.com/JeppeKorgaard98/mAssignment1/tree/ProtoAttempt/protoAttempt" //"/Users/Freja/Documents/Programmering/3.semester/disys/Mandatory1/mAssignment1/protoAttempt"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCourseServer
}

//DATA HERE
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

func (s *server) GetCoursesList(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	var num, err = strconv.Atoi(in.GetName())
	if num <= len(courses) {
		if err != nil {
			fmt.Println("error")
			//error handling
		}
		log.Printf("Received: %v", courses[num])
		return &pb.HelloReply{Message: fmt.Sprint(courses[num].Name) + ";" + fmt.Sprint(courses[num].StudentSatisfactionRating)}, nil
	} else {
		return &pb.HelloReply{Message: "error;error"}, nil
	}
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
