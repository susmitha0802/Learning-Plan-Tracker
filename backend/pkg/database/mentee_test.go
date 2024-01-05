package database

import (
	"errors"
	"lpt/pkg/models"
	reflect "reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	// "gorm.io/driver/postgres"
)

func TestListAssignedCourses(t *testing.T) {
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}

	test_cases := []struct {
		label          string
		request        string
		mockFunc       func()
		expectedOutput []int32
		expectedError  error
	}{
		{
			label:   "Success",
			request: "User Email",
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM "users" WHERE email = $1`)).
					WithArgs("User Email").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses_assigned" WHERE mentee_id = $1 AND "courses_assigned"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"course_id"}).
						AddRow(1).AddRow(2))
			},
			expectedOutput: []int32{1, 2},
			expectedError:  nil,
		},
		{
			label:   "Failure",
			request: "User Email",
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM "users" WHERE email = $1`)).
					WithArgs("User Email").
					WillReturnError(errors.New("User Id not found"))

			},
			expectedOutput: nil,
			expectedError:  errors.New("User Id not found"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.ListAssignedCourses(test_case.request)

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
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}

	course := models.Course{
		Name:    "Course Name",
		Caption: "Course Caption",
		Logo:    "Course Logo",
		Time:    5,
	}

	exercise1 := models.Exercise{}
	exercise1.ID = 1
	exercise1.TopicId = 1
	exercise2 := models.Exercise{}
	exercise2.ID = 2
	exercise2.TopicId = 2

	topic1 := models.Topic{}
	topic1.ID = 1
	topic1.CourseId = 1
	topic1.Exercise = []models.Exercise{exercise1}
	topic2 := models.Topic{}
	topic2.ID = 2
	topic2.CourseId = 1
	topic2.Exercise = []models.Exercise{exercise2}

	course.ID = 1
	course.Topic = []models.Topic{topic1, topic2}

	test_cases := []struct {
		label          string
		request        int32
		mockFunc       func()
		expectedOutput models.Course
		expectedError  error
	}{
		{
			label:   "Success",
			request: 1,
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "caption", "logo", "time"}).
						AddRow(1, "Course Name", "Course Caption", "Course Logo", 5))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "topics" WHERE "topics"."course_id" = $1 AND "topics"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "course_id"}).
						AddRow(1, 1).AddRow(2, 1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."topic_id" IN ($1,$2) AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1, 2).
					WillReturnRows(sqlmock.NewRows([]string{"id", "topic_id"}).
						AddRow(1, 1).AddRow(2, 2))
			},
			expectedOutput: course,
			expectedError:  nil,
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.GetAssignedCourseDetailsByCourseId(test_case.request)

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
func TestGetAssignedMentorDetails(t *testing.T) {
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}

	type request struct {
		course_id    int32
		mentee_email string
	}

	mentor := models.User{
		Name:  "Mentor Name",
		Email: "Mentor Email",
		Role:  1,
	}

	mentor.ID = 1

	test_cases := []struct {
		label          string
		request        request
		mockFunc       func()
		expectedOutput *models.User
		expectedError  error
	}{
		{
			label: "Success",
			request: request{
				course_id:    1,
				mentee_email: "Mentee Email",
			},
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM "users" WHERE email = $1`)).
					WithArgs("Mentee Email").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses_assigned" WHERE (mentee_id = $1 AND course_id = $2) AND "courses_assigned"."deleted_at" IS NULL`)).
					WithArgs(1, 1).
					WillReturnRows(sqlmock.NewRows([]string{"mentor_id"}).
						AddRow(1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role"}).
						AddRow(1, "Mentor Name", "Mentor Email", 1))

			},
			expectedOutput: &mentor,
			expectedError:  nil,
		},
		{
			label: "Failed while getting mentee id",
			request: request{
				course_id:    1,
				mentee_email: "Mentee Email",
			},
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM "users" WHERE email = $1`)).
					WithArgs("Mentee Email").
					WillReturnError(errors.New("User Id not found"))

			},
			expectedOutput: nil,
			expectedError:  errors.New("User Id not found"),
		},
		{
			label: "Failed while getting mentor details",
			request: request{
				course_id:    1,
				mentee_email: "Mentee Email",
			},
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM "users" WHERE email = $1`)).
					WithArgs("Mentee Email").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses_assigned" WHERE (mentee_id = $1 AND course_id = $2) AND "courses_assigned"."deleted_at" IS NULL`)).
					WithArgs(1, 1).
					WillReturnRows(sqlmock.NewRows([]string{"mentor_id"}).
						AddRow(1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnError(errors.New("sql error"))

			},
			expectedOutput: nil,
			expectedError:  errors.New("sql error"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.GetAssignedMentorDetails(test_case.request.course_id, test_case.request.mentee_email)

			if test_case.expectedError != nil {
				if err.Error() != test_case.expectedError.Error() {
					t.Errorf("Error: Expected %v but got %v", test_case.expectedError,
						err)
				}
			}
			// test_case.expectedOutput.ID = 1

			if !reflect.DeepEqual(got, test_case.expectedOutput) {
				t.Errorf("Output: Expected %v but got %v", test_case.expectedOutput,
					got)
			}
		})
	}
}
func TestSubmitExercise(t *testing.T) {
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}
	test_cases := []struct {
		label          string
		request        models.SubmittedExercises
		mockFunc       func()
		expectedOutput int32
		expectedError  error
	}{
		{
			label: "Success",
			request: models.SubmittedExercises{
				MenteeId:   1,
				ExerciseId: 2,
				FileName:   "File Name",
				File:       "File Link",
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "submitted_exercises" ("created_at","updated_at","deleted_at","mentee_id","exercise_id","file_name","file") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1, 2, "File Name", "File Link").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
			expectedOutput: 1,
			expectedError:  nil,
		},
		{
			label: "Failure",
			request: models.SubmittedExercises{
				MenteeId:   1,
				ExerciseId: 2,
				FileName:   "File Name",
				File:       "File Link",
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "submitted_exercises" ("created_at","updated_at","deleted_at","mentee_id","exercise_id","file_name","file") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1, 2, "File Name", "File Link").
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
			got, err := DBClient.SubmitExercise(test_case.request)

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
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}

	type request struct {
		mentee_id   int32
		exercise_id int32
	}

	test_cases := []struct {
		label          string
		request        request
		mockFunc       func()
		expectedOutput string
		expectedError  error
	}{
		{
			label: "Success",
			request: request{
				mentee_id:   1,
				exercise_id: 1,
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "submitted_exercises" SET "deleted_at"=$1 WHERE (mentee_id = $2 AND exercise_id = $3) AND "submitted_exercises"."deleted_at" IS NULL`)).
					WithArgs(sqlmock.AnyArg(), 1, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			expectedOutput: "Deleted successfully",
			expectedError:  nil,
		},
		{
			label: "Failed when trying to delete a non-existing record",
			request: request{
				mentee_id:   1,
				exercise_id: 1,
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "submitted_exercises" SET "deleted_at"=$1 WHERE (mentee_id = $2 AND exercise_id = $3) AND "submitted_exercises"."deleted_at" IS NULL`)).
					WithArgs(sqlmock.AnyArg(), 1, 1).
					WillReturnResult(sqlmock.NewResult(1, 0))

				mock.ExpectCommit()
			},
			expectedOutput: "Already deleted / No record found to delete",
			expectedError:  nil,
		},
		{
			label: "Failed when there is an error",
			request: request{
				mentee_id:   1,
				exercise_id: 1,
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "submitted_exercises" SET "deleted_at"=$1 WHERE (mentee_id = $2 AND exercise_id = $3) AND "submitted_exercises"."deleted_at" IS NULL`)).
					WithArgs(sqlmock.AnyArg(), 1, 1).
					WillReturnError(errors.New("sql error"))

				mock.ExpectRollback()
			},
			expectedOutput: "",
			expectedError:  errors.New("sql error"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.DeleteExercise(test_case.request.mentee_id, test_case.request.exercise_id)

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
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}

	type request struct {
		mentee_id   int32
		exercise_id int32
	}

	type expectedOutput struct {
		file_name string
		file      string
	}

	test_cases := []struct {
		label          string
		request        request
		mockFunc       func()
		expectedOutput expectedOutput
		expectedError  error
	}{
		{
			label: "Success",
			request: request{
				mentee_id:   1,
				exercise_id: 1,
			},
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "submitted_exercises" WHERE (mentee_id = $1 AND exercise_id = $2) AND "submitted_exercises"."deleted_at" IS NULL`)).
					WithArgs(1, 1).
					WillReturnRows(sqlmock.NewRows([]string{"file_name", "file"}).
						AddRow("File Name", "File Link"))
			},
			expectedOutput: expectedOutput{
				file_name: "File Name",
				file:      "File Link",
			},
			expectedError: nil,
		},
		{
			label: "Failure",
			request: request{
				mentee_id:   1,
				exercise_id: 1,
			},
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "submitted_exercises" WHERE (mentee_id = $1 AND exercise_id = $2) AND "submitted_exercises"."deleted_at" IS NULL`)).
					WithArgs(1, 1).
					WillReturnError(errors.New("sql error"))

			},
			expectedOutput: expectedOutput{
				file_name: "",
				file:      "",
			},
			expectedError: errors.New("sql error"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			gotFileName, gotFile, err := DBClient.GetSubmittedExercise(test_case.request.mentee_id, test_case.request.exercise_id)

			if test_case.expectedError != nil {
				if err.Error() != test_case.expectedError.Error() {
					t.Errorf("Error: Expected %v but got %v", test_case.expectedError,
						err)
				}
			}

			if !reflect.DeepEqual(gotFileName, test_case.expectedOutput.file_name) {
				t.Errorf("Output: Expected %v but got %v", test_case.expectedOutput,
					gotFileName)
			}

			if !reflect.DeepEqual(gotFile, test_case.expectedOutput.file) {
				t.Errorf("Output: Expected %v but got %v", test_case.expectedOutput,
					gotFile)
			}
		})
	}
}

func TestListSubmittedExercises(t *testing.T) {
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}

	submittedExercise := models.SubmittedExercises{
		MenteeId:   1,
		ExerciseId: 1,
		FileName:   "File Name",
		File:       "File",
	}

	submittedExercise.Exercise.ID = 1

	test_cases := []struct {
		label          string
		request        int32
		mockFunc       func()
		expectedOutput []models.SubmittedExercises
		expectedError  error
	}{
		{
			label:   "Success",
			request: 1,
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "submitted_exercises" WHERE mentee_id = $1 AND "submitted_exercises"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"mentee_id", "exercise_id", "file_name", "file"}).
						AddRow(1, 1, "File Name", "File"))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."id" = $1 AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))
			},
			expectedOutput: []models.SubmittedExercises{submittedExercise},
			expectedError:  nil,
		},
		{
			label:   "Failed when mentee didnt submit exercise",
			request: 1,
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "submitted_exercises" WHERE mentee_id = $1 AND "submitted_exercises"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{}))

			},
			expectedOutput: nil,
			expectedError:  nil,
		},
		{
			label:   "Failed because of db",
			request: 1,
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "submitted_exercises" WHERE mentee_id = $1 AND "submitted_exercises"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"mentee_id", "exercise_id", "file_name", "file"}).
						AddRow(1, 1, "File Name", "File"))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."id" = $1 AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnError(errors.New("sql error"))

			},
			expectedOutput: nil,
			expectedError:  errors.New("sql error"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.ListSubmittedExercises(test_case.request)

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
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}

	type request struct {
		mentee_id int32
		course_id int32
	}

	test_cases := []struct {
		label          string
		request        request
		mockFunc       func()
		expectedOutput int32
		expectedError  error
	}{
		{
			label: "Success",
			request: request{
				mentee_id: 1,
				course_id: 1,
			},
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "topics" WHERE "topics"."course_id" = $1 AND "topics"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "course_id"}).
						AddRow(1, 1).AddRow(2, 1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."topic_id" IN ($1,$2) AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1, 2).
					WillReturnRows(sqlmock.NewRows([]string{"id", "topic_id"}).
						AddRow(1, 1).AddRow(2, 2))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "submitted_exercises" WHERE mentee_id = $1 AND "submitted_exercises"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "exercise_id"}).
						AddRow(1, 1).AddRow(2, 2))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."id" IN ($1,$2) AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1, 2).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1).AddRow(2))
			},
			expectedOutput: 100,
			expectedError:  nil,
		},
		{
			label: "Failed while getting list of exercise ids",
			request: request{
				mentee_id: 1,
				course_id: 1,
			},
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "topics" WHERE "topics"."course_id" = $1 AND "topics"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "course_id"}).
						AddRow(1, 1).AddRow(2, 1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."topic_id" IN ($1,$2) AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1, 2).
					WillReturnError(errors.New("sql error"))

			},
			expectedOutput: 0,
			expectedError:  errors.New("sql error"),
		},
		{
			label: "Failed while getting submitted exercises",
			request: request{
				mentee_id: 1,
				course_id: 1,
			},
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "topics" WHERE "topics"."course_id" = $1 AND "topics"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "course_id"}).
						AddRow(1, 1).AddRow(2, 1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."topic_id" IN ($1,$2) AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1, 2).
					WillReturnRows(sqlmock.NewRows([]string{"id", "topic_id"}).
						AddRow(1, 1).AddRow(2, 2))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "submitted_exercises" WHERE mentee_id = $1 AND "submitted_exercises"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "exercise_id"}).
						AddRow(1, 1).AddRow(2, 2))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."id" IN ($1,$2) AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1, 2).
					WillReturnError(errors.New("sql error"))

			},
			expectedOutput: 0,
			expectedError:  errors.New("sql error"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.GetProgress(test_case.request.mentee_id, test_case.request.course_id)

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
