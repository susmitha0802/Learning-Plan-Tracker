package main

import (
	"log"

	pb "lpt/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8888"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect the server %v", err)
	}

	defer conn.Close()

	client := pb.NewLearningPlanTrackerServiceClient(conn)

	AddCourse(client)
	AddTopic(client)
	AddExercise(client)
	GetCourses(client)
}
