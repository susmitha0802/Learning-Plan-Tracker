package server

import (
	"context"
	"errors"
	"lpt/pkg/models"
	pb "lpt/pkg/proto"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestListAssignedMenteesAndCourses(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.ListAssignedMenteesAndCoursesRequest
		mockFunc       func()
		expectedOutput *pb.ListAssignedMenteesAndCoursesResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.ListAssignedMenteesAndCoursesRequest{
				MentorEmail: "Mentor Email",
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentor Email").
					Return(int32(1), nil)

				mockDb.EXPECT().ListAssignedMenteesAndCourses(int32(1)).
					Return(
						[]string{"Mentee1 Email", "Mentee2 Email"},
						[]int32{1, 2},
						[]int32{1, 2},
						nil,
					)
			},
			expectedOutput: &pb.ListAssignedMenteesAndCoursesResponse{
				Macd: []*pb.MenteeAndCourseDetails{
					{
						MenteeEmail: "Mentee1 Email",
						MenteeId:    1,
						CourseId:    1,
					},
					{
						MenteeEmail: "Mentee2 Email",
						MenteeId:    2,
						CourseId:    2,
					},
				},
			},
			expectedError: nil,
		},
		{
			label: "Failed while getting mentor id",
			request: &pb.ListAssignedMenteesAndCoursesRequest{
				MentorEmail: "Mentor Email",
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentor Email").
					Return(int32(1),
						errors.New("Unable to get mentor id"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to get mentor id"),
		},
		{
			label: "Failed getting list of assigned mentees and courses",
			request: &pb.ListAssignedMenteesAndCoursesRequest{
				MentorEmail: "Mentor Email",
			},
			mockFunc: func() {
				mockDb.EXPECT().GetUserIdByEmail("Mentor Email").
					Return(int32(1), nil)

				mockDb.EXPECT().ListAssignedMenteesAndCourses(int32(1)).
					Return(nil, nil, nil,
						errors.New("Unable to get list of assigned mentees and courses"),
					)
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to get list of assigned mentees and courses"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.ListAssignedMenteesAndCourses(ctx, test_case.request)

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

func TestListSubmittedExercisesByMentee(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.ListSubmittedExercisesByMenteeRequest
		mockFunc       func()
		expectedOutput *pb.ListSubmittedExercisesByMenteeResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.ListSubmittedExercisesByMenteeRequest{
				MenteeId: 1,
				CourseId: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().
					ListSubmittedExercisesByMentee(int32(1), int32(1)).
					Return([]models.ListSubmittedExercisesByMentee{
						{
							ExerciseId: 1,
							FileName:   "File Name",
							File:       "File Link",
							Question:   "Question",
						},
						{
							ExerciseId: 2,
							FileName:   "File Name",
							File:       "File Link",
							Question:   "Question",
						},
					}, nil)
			},
			expectedOutput: &pb.ListSubmittedExercisesByMenteeResponse{
				SED: []*pb.SubmittedExercisesDetails{
					{
						ExerciseId: 1,
						FileName:   "File Name",
						File:       "File Link",
						Question:   "Question",
					},
					{
						ExerciseId: 2,
						FileName:   "File Name",
						File:       "File Link",
						Question:   "Question",
					},
				},
			},
			expectedError: nil,
		},
		{
			label: "Failure",
			request: &pb.ListSubmittedExercisesByMenteeRequest{
				MenteeId: 1,
				CourseId: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().
					ListSubmittedExercisesByMentee(int32(1), int32(1)).
					Return(nil, errors.New("Unable to list submitted exercises"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Unable to list submitted exercises"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.ListSubmittedExercisesByMentee(ctx, test_case.request)

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
