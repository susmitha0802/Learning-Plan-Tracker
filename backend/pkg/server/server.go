package server

import (
	pb "lpt/pkg/proto"
)

type LearningPlanTrackerServer struct {
	pb.LearningPlanTrackerServiceServer
	DB Database
}
