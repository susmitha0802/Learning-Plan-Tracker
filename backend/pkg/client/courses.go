package main

import (
	"context"
	"log"
	pb "lpt/pkg/proto"
	"time"
)

func AddCourse(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, course := range Courses {

		res, err := client.AddCourse(ctx, &pb.AddCourseRequest{Cd: course})

		if err != nil {
			log.Fatalf("Could not create: %v", err)
		}

		if res.GetCd().Id == 0 {
			log.Fatalf("Not created successfully: %v", err)
		}

		log.Printf("%v Course with id %v created successfully", res.GetCd().Name, res.GetCd().Id)
	}
}

func AddTopic(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, topic := range Topics {
		res, err := client.AddTopic(ctx, &pb.AddTopicRequest{Td: topic})
		if err != nil {
			log.Fatalf("Could not create: %v", err)
		}
		if res.GetTd().Id == 0 {
			log.Fatalf("Not created successfully: %v", err)
		}

		log.Printf("%v Topic with id %v created successfully", res.GetTd().Name, res.GetTd().Id)
	}
}

func AddExercise(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, exercise := range Exercises {
		res, err := client.AddExercise(ctx, &pb.AddExerciseRequest{Ed: exercise})
		if err != nil {
			log.Fatalf("Could not create: %v", err)
		}
		if res.GetEd().Id == 0 {
			log.Fatalf("Not created successfully: %v", err)
		}

		log.Printf("%v Exercise with id %v created successfully", res.GetEd().Question, res.GetEd().Id)
	}
}

func ListCourses(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.ListCourses(ctx, &pb.ListCoursesRequest{})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Println("Done")
}

