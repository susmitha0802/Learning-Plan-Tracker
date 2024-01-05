package main

import (
	"log"

	pb "lpt/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8888"
)

func main() {
	conn, err := grpc.Dial("localhost"+port,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect the server %v", err)
	}

	defer conn.Close()

	client := pb.NewLearningPlanTrackerServiceClient(conn)

	AddCourse(client)
	// AddTopic(client)
	// AddExercise(client)
	// ListCourses(client)
	// GetTotalNoOfExercices(client)
	// AddUser(client)
	// GetUserDetails(client)
	// ListUsersByRole(client)
	// CreateAssignment(client)
	// ListCurrentAssignments(client)
	// ListAssignedCourses(client)
	// GetAssignedCourseDetailsByCourseId(client)
	// GetAssignedCourseAndMentorDetails(client)
	// SubmitExercise(client)
	// DeleteExercise(client)
	// GetSubmittedExercise(client)
	// GetProgress(client)
	// ListAssignedMenteesAndCourses(client)
	// ListSubmittedExercisesByMentee(client)
}
