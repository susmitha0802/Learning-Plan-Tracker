package server

import (
	"context"
	"log"
	pb "lpt/pkg/proto"
)

func (s *LearningPlanTrackerServer) ListAssignedMenteesAndCourses(ctx context.Context, req *pb.ListAssignedMenteesAndCoursesRequest) (*pb.ListAssignedMenteesAndCoursesResponse, error) {

	mentorEmail := req.GetMentorEmail()

	mentorId, err := s.DB.GetUserIdByEmail(mentorEmail)

	if err != nil {
		return nil, err
	}

	log.Println("Get assigned mentee and courses request received")

	menteeEmails, menteeIds, courseIds, err := s.DB.ListAssignedMenteesAndCourses(mentorId)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	var macd []*pb.MenteeAndCourseDetails

	for i := 0; i < len(menteeEmails); i++ {
		macd = append(macd, &pb.MenteeAndCourseDetails{
			MenteeEmail: menteeEmails[i],
			MenteeId:    menteeIds[i],
			CourseId:    courseIds[i],
		})
	}
	return &pb.ListAssignedMenteesAndCoursesResponse{
		Macd: macd,
	}, nil
}

func (s *LearningPlanTrackerServer) ListSubmittedExercisesByMentee(ctx context.Context, req *pb.ListSubmittedExercisesByMenteeRequest) (*pb.ListSubmittedExercisesByMenteeResponse, error) {

	log.Println("Get assigned mentee and courses request received")

	menteeId := req.GetMenteeId()
	courseId := req.GetCourseId()

	submittedExerciseDetails, err := s.DB.ListSubmittedExercisesByMentee(menteeId, courseId)

	if err != nil {
		log.Println("Error", err.Error())
		return nil, err
	}

	var sed []*pb.SubmittedExercisesDetails

	for i := 0; i < len(submittedExerciseDetails); i++ {
		sed = append(sed, &pb.SubmittedExercisesDetails{
			ExerciseId: submittedExerciseDetails[i].ExerciseId,
			FileName:   submittedExerciseDetails[i].FileName,
			File:       submittedExerciseDetails[i].File,
			Question:   submittedExerciseDetails[i].Question,
		})
	}
	return &pb.ListSubmittedExercisesByMenteeResponse{
		SED: sed,
	}, nil
}
