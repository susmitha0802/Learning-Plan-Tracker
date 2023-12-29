package server

import (
	"context"
	"fmt"
	"log"
	"lpt/pkg/alerts"
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

	subject := "Welcome to Learning Plan Tracker!"
	content := fmt.Sprintf("Dear %v\n,We're excited to welcome you to Learning Plan Tracker! Thank you for signing in with name %v and role %v. Happy Learning...", u.Name, u.Name, u.Role)

	alerts.SendEmail(u.Email, subject, content)

	response := pb.User{Id: id, Name: u.Name, Email: u.Email, Role: u.Role}
	return &pb.AddUserResponse{
		User: &response,
	}, nil
}

func (s *LearningPlanTrackerServer) GetUserDetails(ctx context.Context, req *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {
	log.Println("Get user request received")

	id := req.GetId()

	user, err := s.DB.GetUserDetails(id)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	response := pb.User{
		Id:    id,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	return &pb.GetUserDetailsResponse{
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

	courseName, err := s.DB.GetCourseNameById(ca.CourseId)
	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	mentor, err := s.DB.GetUserDetails(ca.MentorId)
	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	mentee, err := s.DB.GetUserDetails(ca.MenteeId)
	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	mentorSubject := fmt.Sprintf("Mentee Assignment Notification for %v", courseName)
	mentorContent := fmt.Sprintf("Dear %v\n,I hope this message finds you well. We are pleased to inform you that you have been selected as the mentor for %v in the upcoming %v course. Your expertise and experience make you an ideal mentor for this course, and we believe your guidance will greatly benefit %v.", mentor.Name, mentee.Name, courseName, mentee.Name)

	alerts.SendEmail(mentor.Email, mentorSubject, mentorContent)

	menteeSubject := fmt.Sprintf("Mentor Assignment Notification for %v", courseName)
	menteeContent := fmt.Sprintf("Dear %v\n,We hope this message finds you well. We are excited to inform you that %v has been assigned as your mentor for the upcoming %v course. %v comes highly recommended and brings valuable expertise to guide you through this learning journey.", mentee.Name, mentor.Name, courseName, mentor.Name)

	alerts.SendEmail(mentee.Email, menteeSubject, menteeContent)

	response := pb.CourseAssignment{
		Id:       id,
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
