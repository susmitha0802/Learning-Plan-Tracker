package main

import (
	"context"
	"log"
	pb "lpt/pkg/proto"
	"time"
)

func AddUser(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, user := range users {

		res, err := client.AddUser(ctx, &pb.AddUserRequest{User: user})

		if err != nil {
			log.Fatalf("Could not create: %v", err)
		}

		if res.GetUser().Id == 0 {
			log.Fatalf("Not created successfully: %v", err)
		}

		log.Printf("%v User with id %v created successfully", res.GetUser().Name, res.GetUser().Id)
	}
}

func GetUsersByRole(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	role := pb.Role_Mentee

	res, err := client.GetUsersByRole(ctx, &pb.GetUsersByRoleRequest{Role: role})

	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}

	log.Printf("There are %v users with role %v ", len(res.GetName()), role)

	for _, v := range res.GetName() {
		log.Printf("%s\n", v)
	}
}

func PostAssignment(client pb.LearningPlanTrackerServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, v := range courses_assigned {

		res, err := client.PostAssignment(ctx, &pb.PostAssignmentRequest{A: v})

		if err != nil {
			log.Fatalf("Could not create: %v", err)
		}

		if res.A.GetId() == 0 {
			log.Fatalf("Not created successfully: %v", err)
		}

		log.Printf("Mentor with id %v is assigned to a mentee with id %v to a %v course successfully", res.A.GetMentorId(), res.A.GetMenteeId(), res.A.GetCourseId())
	}
}