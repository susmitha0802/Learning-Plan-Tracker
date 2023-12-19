package server

import (
	"context"
	"log"
	"lpt/pkg/models"
	pb "lpt/pkg/proto"
)

func (s *LearningPlanTrackerServer) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	u := models.User{
		Name:  req.User.GetName(),
		Email: req.User.GetEmail(),
		Role:  req.User.GetRole(),
	}

	log.Println("Create user request received")

	id, err := s.DB.AddUser(u)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	response := pb.User{Id: int32(id), Name: u.Name, Email: u.Email, Role: u.Role}
	return &pb.AddUserResponse{
		User: &response,
	}, nil
}

func (s *LearningPlanTrackerServer) GetUsersByRole(ctx context.Context, req *pb.GetUsersByRoleRequest) (*pb.GetUsersByRoleResponse, error) {
	role := req.GetRole()

	log.Println("Get users by role request received")

	res, err := s.DB.GetUsersByRole(role)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	return &pb.GetUsersByRoleResponse{
		Name: res,
	}, nil
}
