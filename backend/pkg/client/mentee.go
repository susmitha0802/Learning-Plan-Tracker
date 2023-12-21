package main

import (
	"context"
	"log"
	pb "lpt/pkg/proto"
	"time"
)

func ListAssignedCourses(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userEmail := "susmithapapani@gmail.com"

	res, err := client.ListAssignedCourses(ctx,
		&pb.ListAssignedCoursesRequest{
			Email: userEmail,
		})

	if err != nil {
		log.Fatalf("Could not fetch: %v", err)
	}

	if len(res.GetCourseId()) == 0 {
		log.Println("No courses are assigned")
	} else {
		log.Println("Assigned course id: ")
	}

	for _, v := range res.GetCourseId() {
		log.Printf("%v\n", v)
	}
}

func GetAssignedCourseDetailsByCourseId(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	course_id := 2

	res, err := client.GetAssignedCourseDetailsByCourseId(ctx,
		&pb.GetAssignedCourseDetailsByCourseIdRequest{
			CourseId: int32(course_id)})

	if err != nil {
		log.Fatalf("Could not fetch: %v", err)
	}

	log.Printf("Course details of course id %v:\n", course_id)

	log.Println("Id: ", res.Cd.GetId())
	log.Println("Name: ", res.Cd.GetName())
	log.Println("Caption: ", res.Cd.GetCaption())
	log.Println("Logo: ", res.Cd.GetLogo())
	log.Println("Time: ", res.Cd.GetTime())
	for _, topic := range res.Cd.TD {
		log.Println("\tTopic Id: ", topic.GetId())
		log.Println("\tTopic Name: ", topic.GetName())
		log.Println("\tTopic Resource: ", topic.GetResource())
		for _, exercise := range topic.ED {
			log.Println("\t\tExercise Id: ", exercise.GetId())
			log.Println("\t\tExercise Question: ", exercise.GetQuestion())
		}
	}
}

func GetAssignedCourseAndMentorDetails(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	course_id := 1
	mentee_email := "susmithapapani@gmail.com"

	res, err := client.GetAssignedCourseAndMentorDetails(ctx,
		&pb.GetAssignedCourseAndMentorDetailsRequest{
			CourseId:    int32(course_id),
			MenteeEmail: mentee_email,
		})

	if err != nil {
		log.Fatalf("Could not fetch: %v", err.Error())
	}

	log.Printf("Mentor email %v:\n", res.GetMentorEmail())

	log.Printf("Course details of course id %v:\n", course_id)

	log.Println("Id: ", res.Cd.GetId())
	log.Println("Name: ", res.Cd.GetName())
	log.Println("Caption: ", res.Cd.GetCaption())
	log.Println("Logo: ", res.Cd.GetLogo())
	log.Println("Time: ", res.Cd.GetTime())
	for _, topic := range res.Cd.TD {
		log.Println("\tTopic Id: ", topic.GetId())
		log.Println("\tTopic Name: ", topic.GetName())
		log.Println("\tTopic Resource: ", topic.GetResource())
		for _, exercise := range topic.ED {
			log.Println("\t\tExercise Id: ", exercise.GetId())
			log.Println("\t\tExercise Question: ", exercise.GetQuestion())
		}
	}

}

func SubmitExercise(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, submitExercise := range submitted_exercises {

		res, err := client.SubmitExercise(ctx,
			&pb.SubmitExerciseRequest{
				Sed: submitExercise})

		if err != nil {
			log.Fatalf("Could not submit: %v", err)
		}

		if res.GetId() == 0 {
			log.Fatalf("Not submitted successfully: %v", err)
		}

		log.Printf("Submitted exercise with id %v submitted successfully", res.GetId())
	}
}

func DeleteExercise(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	menteeEmail := "susmithapapani@gmail.com"
	exerciseId := 1

	res, err := client.DeleteExercise(ctx,
		&pb.DeleteExerciseRequest{
			MenteeEmail: menteeEmail,
			ExerciseId:  int32(exerciseId),
		})

	if err != nil {
		log.Fatalf("Could not submit: %v", err)
	}

	log.Printf("%v", res.GetMessage())

}

func GetSubmittedExercise(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	menteeEmail := "mentee3@gmail.com"
	exerciseId := 1

	res, err := client.GetSubmittedExercise(ctx,
		&pb.GetSubmittedExerciseRequest{
			MenteeEmail: menteeEmail,
			ExerciseId:  int32(exerciseId),
		})

	if err != nil {
		log.Fatalf("Could not submit: %v", err)
	}

	if res.GetFileName() == "" {
		log.Println("Exercise is not submitted yet")
	}

	log.Printf("%v\t%v", res.GetFileName(), res.GetFile())
}

func GetProgress(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	menteeEmail := "mentee1@gmail.com"
	courseId := 1

	res, err := client.GetProgress(ctx,
		&pb.GetProgressRequest{
			MenteeEmail: menteeEmail,
			CourseId:    int32(courseId),
		})

	if err != nil {
		log.Fatalf("Could not submit: %v", err)
	}

	log.Printf("Progress is %v", res.GetProgress())

}
