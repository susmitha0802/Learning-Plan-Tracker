package main

import (
	"context"
	"log"
	pb "lpt/pkg/proto"
	"time"
)

func ListAssignedMenteesAndCourses(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	mentorEmail := "susmitha0802@gmail.com"

	res, err := client.ListAssignedMenteesAndCourses(ctx,
		&pb.ListAssignedMenteesAndCoursesRequest{
			MentorEmail: mentorEmail,
		})

	if err != nil {
		log.Fatalf("Could not submit: %v", err)
	}

	if len(res.Macd) == 0 {
		log.Println("No mentee is assigned yet")
	}

	for _, v := range res.Macd {
		log.Printf("Mentee email is %v and course id is %v", v.GetMenteeEmail(), v.GetCourseId())
	}
}

func ListSubmittedExercisesByMentee(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	menteeId := 6
	courseId := 11

	res, err := client.ListSubmittedExercisesByMentee(ctx,
		&pb.ListSubmittedExercisesByMenteeRequest{
			MenteeId: int32(menteeId),
			CourseId: int32(courseId),
		})

	if err != nil {
		log.Fatalf("Could not submit: %v", err)
	}

	if len(res.SED) == 0 {
		log.Println("No exercise is submitted yet")
	}

	for _, v := range res.SED {
		log.Printf("Exercise id: %v", v.GetExerciseId())
		log.Printf("File name: %v", v.GetFileName())
		log.Printf("File: %v", v.GetFile())
		log.Printf("Question: %v", v.GetQuestion())
	}
}
