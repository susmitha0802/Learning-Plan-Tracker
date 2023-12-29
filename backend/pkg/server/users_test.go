package server

import (
	"context"
	"errors"
	"lpt/pkg/alerts"
	"lpt/pkg/database"
	"lpt/pkg/models"
	pb "lpt/pkg/proto"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAddUser(t *testing.T) {
	controller := gomock.NewController(t)
	originalSendEmail := alerts.SendEmail

	defer func() {
		controller.Finish()
		alerts.SendEmail = originalSendEmail
	}()

	mockDb := database.NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	alerts.SendEmail = func(recieverEmail string, subject string, content string) {}

	test_cases := []struct {
		label          string
		request        *pb.AddUserRequest
		mockFunc       func()
		expectedOutput *pb.AddUserResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.AddUserRequest{
				User: &pb.User{
					Name:  "User Name",
					Email: "User Email",
					Role:  0,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().AddUser(gomock.Any()).Return(int32(1), nil)
			},
			expectedOutput: &pb.AddUserResponse{
				User: &pb.User{
					Id:    1,
					Name:  "User Name",
					Email: "User Email",
					Role:  0,
				},
			},
			expectedError: nil,
		},
		{
			label: "Failure",
			request: &pb.AddUserRequest{
				User: &pb.User{
					Name:  "User Name",
					Email: "User Email",
					Role:  0,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().AddUser(gomock.Any()).Return(int32(0),
					errors.New("User can not be added"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("User can not be added"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.AddUser(ctx, test_case.request)

			if test_case.expectedError != nil {
				if err.Error() != test_case.expectedError.Error() {
					t.Errorf("Error: Expected %v but got %v", test_case.expectedError,
						err)
				}
			}

			if !reflect.DeepEqual(got, test_case.expectedOutput) {
				t.Errorf("Output: Expected %v but got %v", test_case.expectedOutput,
					got)
			}
		})
	}
}

func TestGetUserDetails(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := database.NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.GetUserDetailsRequest
		mockFunc       func()
		expectedOutput *pb.GetUserDetailsResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.GetUserDetailsRequest{
				Id: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserDetails(int32(1)).Return(models.User{
					Name:  "User Name",
					Email: "User Email",
					Role:  0,
				}, nil,
				)
			},
			expectedOutput: &pb.GetUserDetailsResponse{
				User: &pb.User{
					Id:    1,
					Name:  "User Name",
					Email: "User Email",
					Role:  0,
				},
			},
			expectedError: nil,
		},
		{
			label: "Failure",
			request: &pb.GetUserDetailsRequest{
				Id: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserDetails(int32(1)).Return(models.User{},
					errors.New("Get user details failed"),
				)
			},
			expectedOutput: nil,
			expectedError:  errors.New("Get user details failed"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.GetUserDetails(ctx, test_case.request)

			if test_case.expectedError != nil {
				if err.Error() != test_case.expectedError.Error() {
					t.Errorf("Error: Expected %v but got %v", test_case.expectedError,
						err)
				}
			}

			if !reflect.DeepEqual(got, test_case.expectedOutput) {
				t.Errorf("Output: Expected %v but got %v", test_case.expectedOutput,
					got)
			}
		})
	}
}

func TestListUsersByRole(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := database.NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.ListUsersByRoleRequest
		mockFunc       func()
		expectedOutput *pb.ListUsersByRoleResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.ListUsersByRoleRequest{
				Role: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().ListUsersByRole(int32(1)).Return([]string{
					"Mentor1",
					"Mentor2",
				}, nil)
			},
			expectedOutput: &pb.ListUsersByRoleResponse{
				Name: []string{
					"Mentor1",
					"Mentor2",
				},
			},
			expectedError: nil,
		},
		{
			label: "Failure",
			request: &pb.ListUsersByRoleRequest{
				Role: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().ListUsersByRole(int32(1)).Return([]string{},
					errors.New("List users by role failed"),
				)
			},
			expectedOutput: nil,
			expectedError:  errors.New("List users by role failed"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.ListUsersByRole(ctx, test_case.request)

			if test_case.expectedError != nil {
				if err.Error() != test_case.expectedError.Error() {
					t.Errorf("Error: Expected %v but got %v", test_case.expectedError,
						err)
				}
			}

			if !reflect.DeepEqual(got, test_case.expectedOutput) {
				t.Errorf("Output: Expected %v but got %v", test_case.expectedOutput,
					got)
			}
		})
	}
}

func TestCreateAssignment(t *testing.T) {
	controller := gomock.NewController(t)
	originalSendEmail := alerts.SendEmail

	defer func() {
		controller.Finish()
		alerts.SendEmail = originalSendEmail
	}()

	mockDb := database.NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	alerts.SendEmail = func(recieverEmail string, subject string, content string) {}

	test_cases := []struct {
		label          string
		request        *pb.CreateAssignmentRequest
		mockFunc       func()
		expectedOutput *pb.CreateAssignmentResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.CreateAssignmentRequest{
				Ca: &pb.CourseAssignment{
					MentorId: 1,
					MenteeId: 2,
					CourseId: 1,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().CreateAssignment(models.CoursesAssigned{
					MentorId: 1,
					MenteeId: 2,
					CourseId: 1,
				}).Return(int32(1), nil)

				mockDb.EXPECT().GetCourseNameById(int32(1)).
					Return("Course Name", nil)

				mockDb.EXPECT().GetUserDetails(int32(1)).
					Return(models.User{
						Name:  "Mentor Name",
						Email: "Mentor Email",
						Role:  1,
					}, nil)

				mockDb.EXPECT().GetUserDetails(int32(2)).
					Return(models.User{
						Name:  "Mentee Name",
						Email: "Mentee Email",
						Role:  1,
					}, nil)
			},
			expectedOutput: &pb.CreateAssignmentResponse{
				Ca: &pb.CourseAssignment{
					Id:       1,
					MentorId: 1,
					MenteeId: 2,
					CourseId: 1,
				},
			},
			expectedError: nil,
		},
		{
			label: "Failed while creating an assignment",
			request: &pb.CreateAssignmentRequest{
				Ca: &pb.CourseAssignment{
					MentorId: 1,
					MenteeId: 2,
					CourseId: 1,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().CreateAssignment(models.CoursesAssigned{
					MentorId: 1,
					MenteeId: 2,
					CourseId: 1,
				}).
					Return(int32(0), errors.New("Assignment is not created"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Assignment is not created"),
		},
		{
			label: "Failed while getting course name",
			request: &pb.CreateAssignmentRequest{
				Ca: &pb.CourseAssignment{
					MentorId: 1,
					MenteeId: 2,
					CourseId: 1,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().CreateAssignment(models.CoursesAssigned{
					MentorId: 1,
					MenteeId: 2,
					CourseId: 1,
				}).Return(int32(1), nil)

				mockDb.EXPECT().GetCourseNameById(int32(1)).
					Return("", errors.New("Unable to get course name"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to get course name"),
		},
		{
			label: "Failed while getting mentor details",
			request: &pb.CreateAssignmentRequest{
				Ca: &pb.CourseAssignment{
					MentorId: 1,
					MenteeId: 2,
					CourseId: 1,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().CreateAssignment(models.CoursesAssigned{
					MentorId: 1,
					MenteeId: 2,
					CourseId: 1,
				}).Return(int32(1), nil)

				mockDb.EXPECT().GetCourseNameById(int32(1)).
					Return("Course Name", nil)

				mockDb.EXPECT().GetUserDetails(int32(1)).
					Return(models.User{}, errors.New("Unable to get mentor details"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to get mentor details"),
		},
		{
			label: "Failed while getting mentee details",
			request: &pb.CreateAssignmentRequest{
				Ca: &pb.CourseAssignment{
					MentorId: 1,
					MenteeId: 2,
					CourseId: 1,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().CreateAssignment(models.CoursesAssigned{
					MentorId: 1,
					MenteeId: 2,
					CourseId: 1,
				}).Return(int32(1), nil)

				mockDb.EXPECT().GetCourseNameById(int32(1)).
					Return("Course Name", nil)

				mockDb.EXPECT().GetUserDetails(int32(1)).
					Return(models.User{
						Name:  "Mentor Name",
						Email: "Mentor Email",
						Role:  1,
					}, nil)

				mockDb.EXPECT().GetUserDetails(int32(2)).
					Return(models.User{}, errors.New("Unable to get mentor details"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to get mentor details"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.CreateAssignment(ctx, test_case.request)

			if test_case.expectedError != nil {
				if err.Error() != test_case.expectedError.Error() {
					t.Errorf("Error: Expected %v but got %v", test_case.expectedError,
						err)
				}
			}

			if !reflect.DeepEqual(got, test_case.expectedOutput) {
				t.Errorf("Output: Expected %v but got %v", test_case.expectedOutput,
					got)
			}
		})
	}
}

func TestListCurrentAssignments(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := database.NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.ListCurrentAssignmentsRequest
		mockFunc       func()
		expectedOutput *pb.ListCurrentAssignmentsResponse
		expectedError  error
	}{
		{
			label:   "Success",
			request: &pb.ListCurrentAssignmentsRequest{},
			mockFunc: func() {
				mockDb.EXPECT().ListCurrentAssignments().
					Return([]models.CoursesAssigned{
						{
							MentorId: 1,
							MenteeId: 2,
							CourseId: 1,
							Mentor: models.User{
								Name: "Mentor 1",
							},
							Mentee: models.User{
								Name: "Mentee 2",
							},
							Course: models.Course{
								Name: "Course 1",
							},
						},
						{
							MentorId: 1,
							MenteeId: 2,
							CourseId: 2,
							Mentor: models.User{
								Name: "Mentor 1",
							},
							Mentee: models.User{
								Name: "Mentee 2",
							},
							Course: models.Course{
								Name: "Course 2",
							},
						},
					}, nil)
			},
			expectedOutput: &pb.ListCurrentAssignmentsResponse{
				Ca: []*pb.CurrentAssignmets{
					{
						MentorName: "Mentor 1",
						MenteeName: "Mentee 2",
						CourseName: "Course 1",
					},
					{
						MentorName: "Mentor 1",
						MenteeName: "Mentee 2",
						CourseName: "Course 2",
					},
				},
			},
			expectedError: nil,
		},
		{
			label:   "Failure",
			request: &pb.ListCurrentAssignmentsRequest{},
			mockFunc: func() {
				mockDb.EXPECT().ListCurrentAssignments().
					Return([]models.CoursesAssigned{},
						errors.New("List current assignments failed"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("List current assignments failed"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.ListCurrentAssignments(ctx, test_case.request)

			if test_case.expectedError != nil {
				if err.Error() != test_case.expectedError.Error() {
					t.Errorf("Error: Expected %v but got %v", test_case.expectedError,
						err)
				}
			}

			if !reflect.DeepEqual(got, test_case.expectedOutput) {
				t.Errorf("Output: Expected %v but got %v", test_case.expectedOutput,
					got)
			}
		})
	}
}
