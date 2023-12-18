package server

import (
	"lpt/pkg/database"
	pb "lpt/pkg/proto"
)

type LearningPlanTrackerServer struct {
	pb.LearningPlanTrackerServiceServer
	DB database.Database
}
