package main

import (
	"log"
	"net"

	"lpt/pkg/database"
	"lpt/pkg/models"
	pb "lpt/pkg/proto"
	"lpt/pkg/server"

	"google.golang.org/grpc"
)

const (
	port = ":8888"
)

func main() {
	db := models.SetUpDB()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterLearningPlanTrackerServiceServer(grpcServer, &server.LearningPlanTrackerServer{
		DB: database.DBClient{DB: db},
	})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the grpc server %v", err)
	}
}
