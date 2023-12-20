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

	response := pb.User{Id: id, Name: u.Name, Email: u.Email, Role: u.Role}
	return &pb.AddUserResponse{
		User: &response,
	}, nil
}

func (s *LearningPlanTrackerServer) ListUsersByRole(ctx context.Context, req *pb.ListUsersByRoleRequest) (*pb.ListUsersByRoleResponse, error) {
	role := req.GetRole()

	log.Println("List users by role request received")

	res, err := s.DB.ListUsersByRole(int32(role.Number()))

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	return &pb.ListUsersByRoleResponse{
		Name: res,
	}, nil
}

func (s *LearningPlanTrackerServer) CreateAssignment(ctx context.Context, req *pb.CreateAssignmentRequest) (*pb.CreateAssignmentResponse, error) {
	ca := models.CoursesAssigned{
		MentorId: req.Ca.GetMentorId(),
		MenteeId: req.Ca.GetMenteeId(),
		CourseId: req.Ca.GetCourseId(),
	}

	log.Println("Create assiggnment request received")

	id, err := s.DB.CreateAssignment(ca)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	response := pb.CourseAssignment{
		Id: id, 
		MentorId: ca.MentorId, 
		MenteeId: ca.MenteeId, 
		CourseId: ca.CourseId,
	}
	return &pb.CreateAssignmentResponse{
		Ca: &response,
	}, nil
}

func (s *LearningPlanTrackerServer) ListCurrentAssignments(ctx context.Context, req *pb.ListCurrentAssignmentsRequest) (*pb.ListCurrentAssignmentsResponse, error) {

	log.Println("Get current assignments request received")

	res, err := s.DB.ListCurrentAssignments()

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	ca := []*pb.CurrentAssignmets{}
	for _, v := range res {
		ca = append(ca, &pb.CurrentAssignmets{
			MentorName: v.Mentor.Name,
			MenteeName: v.Mentee.Name,
			CourseName: v.Course.Name,
		})
	}

	return &pb.ListCurrentAssignmentsResponse{
		Ca: ca,
	}, nil
}
