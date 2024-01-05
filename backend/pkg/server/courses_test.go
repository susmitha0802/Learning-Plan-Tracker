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

func TestAddCourse(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.AddCourseRequest
		mockFunc       func()
		expectedOutput *pb.AddCourseResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.AddCourseRequest{
				Cd: &pb.CourseDetails{
					Name:    "Course Name",
					Caption: "Course Caption",
					Logo:    "Course Logo",
					Time:    5,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().AddCourse(models.Course{
					Name:    "Course Name",
					Caption: "Course Caption",
					Logo:    "Course Logo",
					Time:    5,
				}).Return(int32(1), nil)
			},
			expectedOutput: &pb.AddCourseResponse{
				Cd: &pb.CourseDetails{
					Id:      1,
					Name:    "Course Name",
					Caption: "Course Caption",
					Logo:    "Course Logo",
					Time:    5,
				},
			},
			expectedError: nil,
		},
		{
			label: "Failure",
			request: &pb.AddCourseRequest{
				Cd: &pb.CourseDetails{
					Name:    "Course Name",
					Caption: "Course Caption",
					Logo:    "Course Logo",
					Time:    5,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().AddCourse(models.Course{
					Name:    "Course Name",
					Caption: "Course Caption",
					Logo:    "Course Logo",
					Time:    5,
				}).Return(int32(0),
					errors.New("Course can not be added"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Course can not be added"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.AddCourse(ctx, test_case.request)

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

func TestAddTopic(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.AddTopicRequest
		mockFunc       func()
		expectedOutput *pb.AddTopicResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.AddTopicRequest{
				Td: &pb.TopicDetails{
					Name:     "Topic Name",
					Resource: "Topic Resource",
					CourseId: 1,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().AddTopic(models.Topic{
					Name:     "Topic Name",
					Resource: "Topic Resource",
					CourseId: 1,
				}).Return(int32(1), nil)
			},
			expectedOutput: &pb.AddTopicResponse{
				Td: &pb.TopicDetails{
					Id:       1,
					Name:     "Topic Name",
					Resource: "Topic Resource",
					CourseId: 1,
				},
			},
			expectedError: nil,
		},
		{
			label: "Failure",
			request: &pb.AddTopicRequest{
				Td: &pb.TopicDetails{
					Name:     "Topic Name",
					Resource: "Topic Resource",
					CourseId: 1,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().AddTopic(models.Topic{
					Name:     "Topic Name",
					Resource: "Topic Resource",
					CourseId: 1,
				}).Return(int32(0), errors.New("Topic can not be added"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Topic can not be added"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.AddTopic(ctx, test_case.request)

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

func TestAddExercise(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.AddExerciseRequest
		mockFunc       func()
		expectedOutput *pb.AddExerciseResponse
		expectedError  error
	}{
		{
			label: "Success",
			request: &pb.AddExerciseRequest{
				Ed: &pb.ExerciseDetails{
					Question: "Exercise Question",
					TopicId:  1,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().AddExercise(models.Exercise{
					Question: "Exercise Question",
					TopicId:  1,
				}).Return(int32(1), nil)
			},
			expectedOutput: &pb.AddExerciseResponse{
				Ed: &pb.ExerciseDetails{
					Id:       1,
					Question: "Exercise Question",
					TopicId:  1,
				},
			},
			expectedError: nil,
		},
		{
			label: "Failure",
			request: &pb.AddExerciseRequest{
				Ed: &pb.ExerciseDetails{
					Question: "Exercise Question",
					TopicId:  1,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().AddExercise(models.Exercise{
					Question: "Exercise Question",
					TopicId:  1,
				}).Return(int32(0), errors.New("Exercise can not be added"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("Exercise can not be added"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.AddExercise(ctx, test_case.request)

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

func TestListCourses(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	mockDb := NewMockDatabase(controller)
	mockServer := LearningPlanTrackerServer{
		DB: mockDb,
	}

	ctx := context.Background()

	test_cases := []struct {
		label          string
		request        *pb.ListCoursesRequest
		mockFunc       func()
		expectedOutput *pb.ListCoursesResponse
		expectedError  error
	}{
		{
			label:   "Success",
			request: &pb.ListCoursesRequest{},
			mockFunc: func() {
				mockDb.EXPECT().ListCourses().Return()
			},
			expectedOutput: &pb.ListCoursesResponse{},
			expectedError:  nil,
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := mockServer.ListCourses(ctx, test_case.request)

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
