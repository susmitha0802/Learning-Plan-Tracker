package main

import (
	pb "lpt/pkg/proto"
)

var users = []*pb.User{
	{
		Name:  "Admin",
		Email: "susmitha.papani@beautifulcode.in",
		Role:  pb.Role_Admin,
	},
	{
		Name:  "Mentor1",
		Email: "susmitha0802@gmail.com",
		Role:  pb.Role_Mentor,
	},
	{
		Name:  "Mentor2",
		Email: "mentor2@gmail.com",
		Role:  pb.Role_Mentor,
	},
	{
		Name:  "Mentor3",
		Email: "mentor3@gmail.com",
		Role:  pb.Role_Mentor,
	},
	{
		Name:  "Mentee1",
		Email: "susmithapapani@gmail.com",
		Role:  pb.Role_Mentee,
	},
	{
		Name:  "Mentee2",
		Email: "mentee2@gmail.com",
		Role:  pb.Role_Mentee,
	},
	{
		Name:  "Mentee3",
		Email: "mentee3@gmail.com",
		Role:  pb.Role_Mentee,
	},
}

var courses_assigned = []*pb.CourseAssignment{
	{
		MentorId: 2,
		MenteeId: 5,
		CourseId: 1,
	},
	{
		MentorId: 2,
		MenteeId: 5,
		CourseId: 2,
	},
	{
		MentorId: 3,
		MenteeId: 6,
		CourseId: 1,
	},
}

var submitted_exercises = []*pb.SubmitExerciseDetails{
	{
		MenteeEmail: "susmithapapani@gmail.com",
		ExerciseId:  1,
		FileName:    "exercise-1.txt",
		File:        "data:text/plain;base64,ZmUK",		
	},
	{
		MenteeEmail: "susmithapapani@gmail.com",
		ExerciseId:  45,
		FileName:    "exercise-20.txt",
		File:        "data:text/plain;base64,ZmUK",		
	},
	{
		MenteeEmail: "mentee2@gmail.com",
		ExerciseId:  1,
		FileName:    "exercise-1.txt",
		File:        "data:text/plain;base64,ZmUK",		
	},
}
