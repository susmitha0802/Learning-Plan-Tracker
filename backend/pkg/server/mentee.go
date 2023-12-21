package server

import (
	"context"
	"errors"
	"log"
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

	mentorEmail, err := s.DB.GetAssignedCourseAndMentorDetails(courseId, menteeEmail)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	if mentorEmail == "" {
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
		MentorEmail: mentorEmail,
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
