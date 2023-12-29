package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"lpt/pkg/alerts"
	"lpt/pkg/models"
	pb "lpt/pkg/proto"
)

func (s *LearningPlanTrackerServer) ListAssignedCourses(ctx context.Context, req *pb.ListAssignedCoursesRequest) (*pb.ListAssignedCoursesResponse, error) {

	log.Println("Get assigned course id request received")

	userEmail := req.GetEmail()

	res, err := s.DB.ListAssignedCourses(userEmail)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	return &pb.ListAssignedCoursesResponse{
		CourseId: res,
	}, nil
}

func (s *LearningPlanTrackerServer) GetAssignedCourseDetailsByCourseId(ctx context.Context, req *pb.GetAssignedCourseDetailsByCourseIdRequest) (*pb.GetAssignedCourseDetailsByCourseIdResponse, error) {

	log.Println("Get assigned course id request received")

	course_id := req.GetCourseId()

	res, err := s.DB.GetAssignedCourseDetailsByCourseId(course_id)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	td := []*pb.TopicDetails{}

	for _, topic := range res.Topic {

		ed := []*pb.ExerciseDetails{}

		for _, exercise := range topic.Exercise {
			ed = append(ed, &pb.ExerciseDetails{
				Id:       int32(exercise.ID),
				Question: exercise.Question,
			})
		}

		td = append(td, &pb.TopicDetails{
			Id:       int32(topic.ID),
			Name:     topic.Name,
			Resource: topic.Resource,
			ED:       ed,
		})
	}

	cd := pb.CourseDetails{
		Id:      int32(res.ID),
		Name:    res.Name,
		Caption: res.Caption,
		Logo:    res.Logo,
		Time:    res.Time,
		TD:      td,
	}

	return &pb.GetAssignedCourseDetailsByCourseIdResponse{
		Cd: &cd,
	}, nil
}

func (s *LearningPlanTrackerServer) GetAssignedCourseAndMentorDetails(ctx context.Context, req *pb.GetAssignedCourseAndMentorDetailsRequest) (*pb.GetAssignedCourseAndMentorDetailsResponse, error) {

	log.Println("Get assigned course and mentor details request received")

	courseId := req.GetCourseId()
	menteeEmail := req.GetMenteeEmail()

	mentor, err := s.DB.GetAssignedMentorDetails(courseId, menteeEmail)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	if mentor == nil {
		log.Println("Assignment not found")
		return nil, errors.New("Assignment not found")
	}

	courseDetails, err := s.DB.GetAssignedCourseDetailsByCourseId(courseId)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	td := []*pb.TopicDetails{}

	for _, topic := range courseDetails.Topic {

		ed := []*pb.ExerciseDetails{}

		for _, exercise := range topic.Exercise {
			ed = append(ed, &pb.ExerciseDetails{
				Id:       int32(exercise.ID),
				Question: exercise.Question,
			})
		}

		td = append(td, &pb.TopicDetails{
			Id:       int32(topic.ID),
			Name:     topic.Name,
			Resource: topic.Resource,
			ED:       ed,
		})
	}

	cd := pb.CourseDetails{
		Id:      int32(courseDetails.ID),
		Name:    courseDetails.Name,
		Caption: courseDetails.Caption,
		Logo:    courseDetails.Logo,
		Time:    courseDetails.Time,
		TD:      td,
	}

	return &pb.GetAssignedCourseAndMentorDetailsResponse{
		Cd:          &cd,
		MentorEmail: mentor.Email,
	}, nil
}

func (s *LearningPlanTrackerServer) SubmitExercise(ctx context.Context, req *pb.SubmitExerciseRequest) (*pb.SubmitExerciseResponse, error) {

	MenteeEmail := req.Sed.GetMenteeEmail()
	MenteeId, err := s.DB.GetUserIdByEmail(MenteeEmail)

	if err != nil {
		return nil, errors.New("Mentee Id not found")
	}

	submit_exercise := models.SubmittedExercises{
		MenteeId:   MenteeId,
		ExerciseId: req.Sed.GetExerciseId(),
		FileName:   req.Sed.GetFileName(),
		File:       req.Sed.GetFile(),
	}

	log.Println("Submit exercise request received")

	id, err := s.DB.SubmitExercise(submit_exercise)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	courseId := 1
	mentee, err := s.DB.GetUserDetails(submit_exercise.MenteeId)
	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	mentor, err := s.DB.GetAssignedMentorDetails(int32(courseId), mentee.Email)
	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}
	
	subject := fmt.Sprintf("New Exercise Submission from %v", mentee.Name)
	content := fmt.Sprintf("Dear %v, I hope this message finds you well. We wanted to inform you that %v has submitted a new exercise.Please take a moment to review the submission and provide feedback to support %v's learning journey. ", mentor.Name, mentee.Name, mentee.Name)

	alerts.SendEmail(mentor.Email, subject, content)

	return &pb.SubmitExerciseResponse{
		Id: id,
	}, nil
}

func (s *LearningPlanTrackerServer) DeleteExercise(ctx context.Context, req *pb.DeleteExerciseRequest) (*pb.DeleteExerciseResponse, error) {

	MenteeEmail := req.GetMenteeEmail()
	ExerciseId := req.GetExerciseId()

	MenteeId, err := s.DB.GetUserIdByEmail(MenteeEmail)

	if err != nil {
		return nil, errors.New("Mentee Id not found")
	}

	log.Println("Delete exercise request received")

	res, err := s.DB.DeleteExercise(MenteeId, ExerciseId)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	return &pb.DeleteExerciseResponse{
		Message: res,
	}, nil
}

func (s *LearningPlanTrackerServer) GetSubmittedExercise(ctx context.Context, req *pb.GetSubmittedExerciseRequest) (*pb.GetSubmittedExerciseResponse, error) {

	MenteeEmail := req.GetMenteeEmail()
	ExerciseId := req.GetExerciseId()

	MenteeId, err := s.DB.GetUserIdByEmail(MenteeEmail)

	if err != nil {
		return nil, errors.New("Mentee Id not found")
	}

	log.Println("Get submitted exercise request received")

	fileName, file, err := s.DB.GetSubmittedExercise(MenteeId, ExerciseId)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	return &pb.GetSubmittedExerciseResponse{
		FileName: fileName,
		File:     file,
	}, nil
}

func (s *LearningPlanTrackerServer) GetProgress(ctx context.Context, req *pb.GetProgressRequest) (*pb.GetProgressResponse, error) {

	menteeEmail := req.GetMenteeEmail()
	courseId := req.GetCourseId()

	menteeId, err := s.DB.GetUserIdByEmail(menteeEmail)

	if err != nil {
		return nil, errors.New("Mentee Id not found")
	}

	log.Println("Get progress request received")

	res, err := s.DB.GetProgress(menteeId, courseId)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	return &pb.GetProgressResponse{
		Progress: res,
	}, nil
}

