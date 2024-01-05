package server

import (
	"context"
	"errors"
	"lpt/pkg/alerts"
	"lpt/pkg/models"
	pb "lpt/pkg/proto"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestListAssignedCourses(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.ListAssignedCoursesRequest
		mockFunc       func()
		expectedOutput *pb.ListAssignedCoursesResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.ListAssignedCoursesRequest{
				Email: "Email",
			},
			mockFunc: func() {
				mockDb.EXPECT().ListAssignedCourses("Email").
					Return([]int32{1, 2}, nil)
			},
			expectedOutput: &pb.ListAssignedCoursesResponse{
				CourseId: []int32{1, 2},
			},
			expectedError: nil,
		},
		{
			label: "Failure",
			request: &pb.ListAssignedCoursesRequest{
				Email: "Email",
			},
			mockFunc: func() {
				mockDb.EXPECT().ListAssignedCourses("Email").
					Return(nil,
						errors.New("Unable to list assigned courses"),
					)
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to list assigned courses"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.ListAssignedCourses(ctx, test_case.request)

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

func TestGetAssignedCourseDetailsByCourseId(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.GetAssignedCourseDetailsByCourseIdRequest
		mockFunc       func()
		expectedOutput *pb.GetAssignedCourseDetailsByCourseIdResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.GetAssignedCourseDetailsByCourseIdRequest{
				CourseId: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetAssignedCourseDetailsByCourseId(int32(1)).
					Return(models.Course{
						Name:    "Course Name",
						Caption: "Course Caption",
						Logo:    "Course Logo",
						Time:    1,
						Topic: []models.Topic{
							{
								Name:     "Topic1 Name",
								Resource: "Topic1 Resource",
								Exercise: []models.Exercise{
									{
										Question: "Exercise1 Question",
									},
									{
										Question: "Exercise2 Question",
									},
								},
							},
						},
					}, nil)
			},
			expectedOutput: &pb.GetAssignedCourseDetailsByCourseIdResponse{
				Cd: &pb.CourseDetails{
					Name:    "Course Name",
					Caption: "Course Caption",
					Logo:    "Course Logo",
					Time:    1,
					TD: []*pb.TopicDetails{
						{
							Name:     "Topic1 Name",
							Resource: "Topic1 Resource",
							ED: []*pb.ExerciseDetails{
								{
									Question: "Exercise1 Question",
								},
								{
									Question: "Exercise2 Question",
								},
							},
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			label: "Failure",
			request: &pb.GetAssignedCourseDetailsByCourseIdRequest{
				CourseId: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetAssignedCourseDetailsByCourseId(int32(1)).
					Return(models.Course{},
						errors.New("Unable to get assigned course details"),
					)
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to get assigned course details"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.GetAssignedCourseDetailsByCourseId(ctx, test_case.request)

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

func TestGetAssignedCourseAndMentorDetails(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.GetAssignedCourseAndMentorDetailsRequest
		mockFunc       func()
		expectedOutput *pb.GetAssignedCourseAndMentorDetailsResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.GetAssignedCourseAndMentorDetailsRequest{
				CourseId:    1,
				MenteeEmail: "Mentee Email",
			},
			mockFunc: func() {
				mockDb.EXPECT().
					GetAssignedMentorDetails(int32(1), "Mentee Email").
					Return(&models.User{
						Name:  "Mentor Name",
						Email: "Mentor Email",
						Role:  1,
					}, nil)
				mockDb.EXPECT().GetAssignedCourseDetailsByCourseId(int32(1)).
					Return(models.Course{
						Name:    "Course Name",
						Caption: "Course Caption",
						Logo:    "Course Logo",
						Time:    1,
						Topic: []models.Topic{
							{
								Name:     "Topic1 Name",
								Resource: "Topic1 Resource",
								Exercise: []models.Exercise{
									{
										Question: "Exercise1 Question",
									},
									{
										Question: "Exercise2 Question",
									},
								},
							},
						},
					}, nil)
			},
			expectedOutput: &pb.GetAssignedCourseAndMentorDetailsResponse{
				Cd: &pb.CourseDetails{
					Name:    "Course Name",
					Caption: "Course Caption",
					Logo:    "Course Logo",
					Time:    1,
					TD: []*pb.TopicDetails{
						{
							Name:     "Topic1 Name",
							Resource: "Topic1 Resource",
							ED: []*pb.ExerciseDetails{
								{
									Question: "Exercise1 Question",
								},
								{
									Question: "Exercise2 Question",
								},
							},
						},
					},
				},
				MentorEmail: "Mentor Email",
			},
			expectedError: nil,
		},
		{
			label: "Failed while getting mentor details",
			request: &pb.GetAssignedCourseAndMentorDetailsRequest{
				CourseId:    1,
				MenteeEmail: "Mentee Email",
			},
			mockFunc: func() {
				mockDb.EXPECT().
					GetAssignedMentorDetails(int32(1), "Mentee Email").
					Return(&models.User{},
						errors.New("Unable to get assigned mentor details"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to get assigned mentor details"),
		},
		{
			label: "Failed while getting course details",
			request: &pb.GetAssignedCourseAndMentorDetailsRequest{
				CourseId:    1,
				MenteeEmail: "Mentee Email",
			},
			mockFunc: func() {
				mockDb.EXPECT().
					GetAssignedMentorDetails(int32(1), "Mentee Email").
					Return(&models.User{
						Name:  "Mentor Name",
						Email: "Mentor Email",
						Role:  1,
					}, nil)
				mockDb.EXPECT().GetAssignedCourseDetailsByCourseId(int32(1)).
					Return(models.Course{},
						errors.New("Unable to get assigned course details"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to get assigned course details"),
		},
		{
			label: "Failed when there is no assignment",
			request: &pb.GetAssignedCourseAndMentorDetailsRequest{
				CourseId:    1,
				MenteeEmail: "Mentee Email",
			},
			mockFunc: func() {
				mockDb.EXPECT().
					GetAssignedMentorDetails(int32(1), "Mentee Email").
					Return(nil, nil)
			},
			expectedOutput: nil,
			expectedError:  errors.New("Assignment not found"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.GetAssignedCourseAndMentorDetails(ctx, test_case.request)

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

func TestSubmitExercise(t *testing.T) {
	controller := gomock.NewController(t)
	originalSendEmail := alerts.SendEmail

	defer func() {
		controller.Finish()
		alerts.SendEmail = originalSendEmail
	}()

	mockDb := NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	alerts.SendEmail = func(recieverEmail string, subject string, content string) {}

	test_cases := []struct {
		label          string
		request        *pb.SubmitExerciseRequest
		mockFunc       func()
		expectedOutput *pb.SubmitExerciseResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.SubmitExerciseRequest{
				Sed: &pb.SubmitExerciseDetails{
					MenteeEmail: "Mentee Email",
					ExerciseId:  1,
					FileName:    "File Name",
					File:        "File Link",
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(1), nil)

				mockDb.EXPECT().SubmitExercise(models.SubmittedExercises{
					MenteeId:   1,
					ExerciseId: 1,
					FileName:   "File Name",
					File:       "File Link",
				}).Return(int32(1), nil)

				mockDb.EXPECT().GetUserDetails(int32(1)).
					Return(models.User{
						Name:  "Mentee Name",
						Email: "Mentee Email",
						Role:  2,
					}, nil)

				mockDb.EXPECT().
					GetAssignedMentorDetails(int32(1), "Mentee Email").
					Return(&models.User{
						Name:  "Mentor Name",
						Email: "Mentor Email",
						Role:  1,
					}, nil)
			},
			expectedOutput: &pb.SubmitExerciseResponse{
				Id: 1,
			},
			expectedError: nil,
		},
		{
			label: "Failed while getting mentee id",
			request: &pb.SubmitExerciseRequest{
				Sed: &pb.SubmitExerciseDetails{
					MenteeEmail: "Mentee Email",
					ExerciseId:  1,
					FileName:    "File Name",
					File:        "File Link",
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(0),
						errors.New("Mentee Id not found"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Mentee Id not found"),
		},
		{
			label: "Failed while submitting the exercise",
			request: &pb.SubmitExerciseRequest{
				Sed: &pb.SubmitExerciseDetails{
					MenteeEmail: "Mentee Email",
					ExerciseId:  1,
					FileName:    "File Name",
					File:        "File Link",
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(1), nil)

				mockDb.EXPECT().SubmitExercise(models.SubmittedExercises{
					MenteeId:   1,
					ExerciseId: 1,
					FileName:   "File Name",
					File:       "File Link",
				}).Return(int32(0),
					errors.New("Unable to submit exercise"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to submit exercise"),
		},
		{
			label: "Failed while getting mentee details",
			request: &pb.SubmitExerciseRequest{
				Sed: &pb.SubmitExerciseDetails{
					MenteeEmail: "Mentee Email",
					ExerciseId:  1,
					FileName:    "File Name",
					File:        "File Link",
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(1), nil)

				mockDb.EXPECT().SubmitExercise(models.SubmittedExercises{
					MenteeId:   1,
					ExerciseId: 1,
					FileName:   "File Name",
					File:       "File Link",
				}).Return(int32(1), nil)

				mockDb.EXPECT().GetUserDetails(int32(1)).
					Return(models.User{},
						errors.New("Unable to get mentee details"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to get mentee details"),
		},
		{
			label: "Failed while getting assigned mentor details",
			request: &pb.SubmitExerciseRequest{
				Sed: &pb.SubmitExerciseDetails{
					MenteeEmail: "Mentee Email",
					ExerciseId:  1,
					FileName:    "File Name",
					File:        "File Link",
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(1), nil)

				mockDb.EXPECT().SubmitExercise(models.SubmittedExercises{
					MenteeId:   1,
					ExerciseId: 1,
					FileName:   "File Name",
					File:       "File Link",
				}).Return(int32(1), nil)

				mockDb.EXPECT().GetUserDetails(int32(1)).
					Return(models.User{
						Name:  "Mentee Name",
						Email: "Mentee Email",
						Role:  2,
					}, nil)

				mockDb.EXPECT().
					GetAssignedMentorDetails(int32(1), "Mentee Email").
					Return(&models.User{},
						errors.New("Unable to get assigned mentor details"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to get assigned mentor details"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.SubmitExercise(ctx, test_case.request)

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

func TestDeleteExercise(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.DeleteExerciseRequest
		mockFunc       func()
		expectedOutput *pb.DeleteExerciseResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.DeleteExerciseRequest{
				MenteeEmail: "Mentee Email",
				ExerciseId:  1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(1), nil)

				mockDb.EXPECT().DeleteExercise(int32(1), int32(1)).
					Return("Deleted successfully", nil)
			},
			expectedOutput: &pb.DeleteExerciseResponse{
				Message: "Deleted successfully",
			},
			expectedError: nil,
		},
		{
			label: "Failed to get mentee id",
			request: &pb.DeleteExerciseRequest{
				MenteeEmail: "Mentee Email",
				ExerciseId:  1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(0),
						errors.New("Mentee Id not found"),
					)
			},
			expectedOutput: nil,
			expectedError:  errors.New("Mentee Id not found"),
		},
		{
			label: "Failed to delete (Already deleted / No record found to delete) exercise",
			request: &pb.DeleteExerciseRequest{
				MenteeEmail: "Mentee Email",
				ExerciseId:  1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(1), nil)

				mockDb.EXPECT().DeleteExercise(int32(1), int32(1)).
					Return("Already deleted / No record found to delete", nil)
			},
			expectedOutput: &pb.DeleteExerciseResponse{
				Message: "Already deleted / No record found to delete",
			},
			expectedError: nil,
		},
		{
			label: "Failed to delete exercise",
			request: &pb.DeleteExerciseRequest{
				MenteeEmail: "Mentee Email",
				ExerciseId:  1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(1), nil)

				mockDb.EXPECT().DeleteExercise(int32(1), int32(1)).
					Return("", errors.New("Unable to delete exercise"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to delete exercise"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.DeleteExercise(ctx, test_case.request)

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

func TestGetSubmittedExercise(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.GetSubmittedExerciseRequest
		mockFunc       func()
		expectedOutput *pb.GetSubmittedExerciseResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.GetSubmittedExerciseRequest{
				MenteeEmail: "Mentee Email",
				ExerciseId:  1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(1), nil)

				mockDb.EXPECT().GetSubmittedExercise(int32(1), int32(1)).
					Return("File Name", "File Link", nil)
			},
			expectedOutput: &pb.GetSubmittedExerciseResponse{
				FileName: "File Name",
				File:     "File Link",
			},
			expectedError: nil,
		},
		{
			label: "Failed to get mentee id",
			request: &pb.GetSubmittedExerciseRequest{
				MenteeEmail: "Mentee Email",
				ExerciseId:  1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(0),
						errors.New("Mentee Id not found"),
					)
			},
			expectedOutput: nil,
			expectedError:  errors.New("Mentee Id not found"),
		},
		{
			label: "Failed to get submitted exercise",
			request: &pb.GetSubmittedExerciseRequest{
				MenteeEmail: "Mentee Email",
				ExerciseId:  1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(1), nil)

				mockDb.EXPECT().GetSubmittedExercise(int32(1), int32(1)).
					Return("", "",
						errors.New("Unable to get submitted exercise"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to get submitted exercise"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.GetSubmittedExercise(ctx, test_case.request)

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

func TestGetProgress(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.GetProgressRequest
		mockFunc       func()
		expectedOutput *pb.GetProgressResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.GetProgressRequest{
				MenteeEmail: "Mentee Email",
				CourseId:    1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(1), nil)

				mockDb.EXPECT().GetProgress(int32(1), int32(1)).
					Return(int32(100), nil)
			},
			expectedOutput: &pb.GetProgressResponse{
				Progress: int32(100),
			},
			expectedError: nil,
		},
		{
			label: "Failed to get mentee id",
			request: &pb.GetProgressRequest{
				MenteeEmail: "Mentee Email",
				CourseId:    1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(0),
						errors.New("Mentee Id not found"),
					)
			},
			expectedOutput: nil,
			expectedError:  errors.New("Mentee Id not found"),
		},
		{
			label: "Failed to get progress",
			request: &pb.GetProgressRequest{
				MenteeEmail: "Mentee Email",
				CourseId:    1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentee Email").
					Return(int32(1), nil)

				mockDb.EXPECT().GetProgress(int32(1), int32(1)).
					Return(int32(0),
						errors.New("Unable to get progress"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to get progress"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.GetProgress(ctx, test_case.request)

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
