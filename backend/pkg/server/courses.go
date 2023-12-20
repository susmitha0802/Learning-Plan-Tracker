package server

import (
	"context"
	"log"
	"lpt/pkg/models"
	pb "lpt/pkg/proto"
)

func (s *LearningPlanTrackerServer) AddCourse(ctx context.Context, req *pb.AddCourseRequest) (*pb.AddCourseResponse, error) {
	c := models.Course{
		Name:    req.Cd.GetName(),
		Caption: req.Cd.GetCaption(),
		Logo:    req.Cd.GetLogo(),
		Time:    req.Cd.GetTime(),
	}

	log.Println("Create course request received")

	id, err := s.DB.AddCourse(c)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	response := pb.CourseDetails{Id: id, Name: c.Name, Caption: c.Caption, Logo: c.Logo, Time: c.Time}
	return &pb.AddCourseResponse{
		Cd: &response,
	}, nil
}

func (s *LearningPlanTrackerServer) AddTopic(ctx context.Context, req *pb.AddTopicRequest) (*pb.AddTopicResponse, error) {
	t := models.Topic{
		Name:     req.Td.GetName(),
		Resource: req.Td.GetResource(),
		CourseId: req.Td.GetCourseId(),
	}

	log.Println("Create topic request received")

	id, err := s.DB.AddTopic(t)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	response := pb.TopicDetails{Id: id, Name: t.Name, Resource: t.Resource, CourseId: t.CourseId}
	return &pb.AddTopicResponse{
		Td: &response,
	}, nil
}

func (s *LearningPlanTrackerServer) AddExercise(ctx context.Context, req *pb.AddExerciseRequest) (*pb.AddExerciseResponse, error) {
	e := models.Exercise{
		Question: req.Ed.GetQuestion(),
		TopicId:  req.Ed.GetTopicId(),
	}

	log.Println("Create exercise request received")

	id, err := s.DB.AddExercise(e)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	response := pb.ExerciseDetails{Id: id, Question: e.Question, TopicId: e.TopicId}
	return &pb.AddExerciseResponse{
		Ed: &response,
	}, nil
}

func (s *LearningPlanTrackerServer) ListCourses(ctx context.Context, req *pb.ListCoursesRequest) (*pb.ListCoursesResponse, error) {

	log.Println("Read courses request received")

	s.DB.ListCourses()

	return &pb.ListCoursesResponse{}, nil
}
