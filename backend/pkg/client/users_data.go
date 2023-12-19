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

