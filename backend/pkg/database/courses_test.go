package database

import (
	"errors"
	"lpt/pkg/models"
	reflect "reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddCourse(t *testing.T) {
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}
	test_cases := []struct {
		label          string
		request        models.Course
		mockFunc       func()
		expectedOutput int32
		expectedError  error
	}{
		{
			label: "Success",
			request: models.Course{
				Name:    "Course Name",
				Caption: "Course Caption",
				Logo:    "Course Logo",
				Time:    5,
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "courses" ("created_at","updated_at","deleted_at","name","caption","logo","time") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Course Name", "Course Caption", "Course Logo", 5).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
			expectedOutput: 1,
			expectedError:  nil,
		},
		{
			label: "Failure",
			request: models.Course{
				Name:    "Course Name",
				Caption: "Course Caption",
				Logo:    "Course Logo",
				Time:    5,
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "courses" ("created_at","updated_at","deleted_at","name","caption","logo","time") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Course Name", "Course Caption", "Course Logo", 5).
					WillReturnError(errors.New("sql error"))

				mock.ExpectRollback()
			},
			expectedOutput: 0,
			expectedError:  errors.New("sql error"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.AddCourse(test_case.request)

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

func TestGetCourseNameById(t *testing.T) {
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}
	test_cases := []struct {
		label          string
		request        int32
		mockFunc       func()
		expectedOutput string
		expectedError  error
	}{
		{
			label:   "Success",
			request: 1,
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT name FROM "courses" WHERE id = $1`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"name"}).
						AddRow("Course Name"))
			},
			expectedOutput: "Course Name",
			expectedError:  nil,
		},
		{
			label:   "Failure",
			request: 1,
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT name FROM "courses" WHERE id = $1`)).
					WithArgs(1).
					WillReturnError(errors.New("sql error"))

			},
			expectedOutput: "",
			expectedError:  errors.New("sql error"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.GetCourseNameById(test_case.request)

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
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}
	test_cases := []struct {
		label          string
		request        models.Topic
		mockFunc       func()
		expectedOutput int32
		expectedError  error
	}{
		{
			label: "Success",
			request: models.Topic{
				Name:     "Topic Name",
				Resource: "Topic Resource",
				CourseId: 1,
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "topics" ("created_at","updated_at","deleted_at","name","resource","course_id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Topic Name", "Topic Resource", 1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
			expectedOutput: 1,
			expectedError:  nil,
		},
		{
			label: "Failure",
			request: models.Topic{
				Name:     "Topic Name",
				Resource: "Topic Resource",
				CourseId: 1,
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "topics" ("created_at","updated_at","deleted_at","name","resource","course_id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Topic Name", "Topic Resource", 1).
					WillReturnError(errors.New("sql error"))

				mock.ExpectRollback()
			},
			expectedOutput: 0,
			expectedError:  errors.New("sql error"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.AddTopic(test_case.request)

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
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}
	test_cases := []struct {
		label          string
		request        models.Exercise
		mockFunc       func()
		expectedOutput int32
		expectedError  error
	}{
		{
			label: "Success",
			request: models.Exercise{
				Question: "Exercise Question",
				TopicId:  1,
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "exercises" ("created_at","updated_at","deleted_at","question","topic_id") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Exercise Question", 1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
			expectedOutput: 1,
			expectedError:  nil,
		},
		{
			label: "Failure",
			request: models.Exercise{
				Question: "Exercise Question",
				TopicId:  1,
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "exercises" ("created_at","updated_at","deleted_at","question","topic_id") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Exercise Question", 1).
					WillReturnError(errors.New("sql error"))

				mock.ExpectRollback()
			},
			expectedOutput: 0,
			expectedError:  errors.New("sql error"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.AddExercise(test_case.request)

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

func TestListExerciseIds(t *testing.T) {
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}
	test_cases := []struct {
		label          string
		request        int32
		mockFunc       func()
		expectedOutput []int32
		expectedError  error
	}{
		{
			label:   "Success",
			request: 1,
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "topics" WHERE "topics"."course_id" = $1 AND "topics"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "course_id"}).AddRow(1, 1).AddRow(2, 1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."topic_id" IN ($1,$2) AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1, 2).
					WillReturnRows(sqlmock.NewRows([]string{"id", "topic_id"}).AddRow(1, 1).AddRow(2, 2))
			},

			expectedOutput: []int32{1, 2},
			expectedError:  nil,
		},
		{
			label:   "Failure",
			request: 1,
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "topics" WHERE "topics"."course_id" = $1 AND "topics"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "course_id"}).AddRow(1, 1).AddRow(2, 1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."topic_id" IN ($1,$2) AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1, 2).
					WillReturnError(errors.New("sql error"))

			},
			expectedOutput: nil,
			expectedError:  errors.New("sql error"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.ListExerciseIds(test_case.request)

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
