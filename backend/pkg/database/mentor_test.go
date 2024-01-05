package database

import (
	"errors"
	"lpt/pkg/models"
	reflect "reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestListAssignedMenteesAndCourses(t *testing.T) {
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}

	type expectedOutput struct {
		mentee_emails []string
		mentee_ids    []int32
		course_ids    []int32
	}

	test_cases := []struct {
		label          string
		request        int32
		mockFunc       func()
		expectedOutput expectedOutput
		expectedError  error
	}{
		{
			label:   "Success",
			request: 1,
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses_assigned" WHERE mentor_id = $1 AND "courses_assigned"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"mentee_id", "course_id"}).
						AddRow(1, 1).AddRow(2, 2))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT email FROM "users" WHERE id = $1`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"email"}).
						AddRow("Mentee1 Email"))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT email FROM "users" WHERE id = $1`)).
					WithArgs(2).
					WillReturnRows(sqlmock.NewRows([]string{"email"}).
						AddRow("Mentee2 Email"))
			},
			expectedOutput: expectedOutput{
				mentee_emails: []string{"Mentee1 Email", "Mentee2 Email"},
				mentee_ids:    []int32{1, 2},
				course_ids:    []int32{1, 2},
			},
			expectedError: nil,
		},
		{
			label:   "Failed while getting courses assigned",
			request: 1,
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses_assigned" WHERE mentor_id = $1 AND "courses_assigned"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnError(errors.New("sql error"))

			},
			expectedOutput: expectedOutput{
				mentee_emails: nil,
				mentee_ids:    nil,
				course_ids:    nil,
			},
			expectedError: errors.New("sql error"),
		},
		{
			label:   "Failed while getting mentee email",
			request: 1,
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses_assigned" WHERE mentor_id = $1 AND "courses_assigned"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"mentee_id", "course_id"}).
						AddRow(1, 1).AddRow(2, 2))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT email FROM "users" WHERE id = $1`)).
					WithArgs(1).
					WillReturnError(errors.New("sql error"))

			},
			expectedOutput: expectedOutput{
				mentee_emails: nil,
				mentee_ids:    nil,
				course_ids:    nil,
			},
			expectedError: errors.New("sql error"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			gotMenteeEmails, gotMenteeIds, gotCourseIds, err := DBClient.ListAssignedMenteesAndCourses(test_case.request)

			if test_case.expectedError != nil {
				if err.Error() != test_case.expectedError.Error() {
					t.Errorf("Error: Expected %v but got %v", test_case.expectedError,
						err)
				}
			}

			if !reflect.DeepEqual(gotMenteeEmails, test_case.expectedOutput.mentee_emails) {
				t.Errorf("Output: Expected %v but got %v", test_case.expectedOutput,
					gotMenteeEmails)
			}

			if !reflect.DeepEqual(gotMenteeIds, test_case.expectedOutput.mentee_ids) {
				t.Errorf("Output: Expected %v but got %v", test_case.expectedOutput,
					gotMenteeIds)
			}

			if !reflect.DeepEqual(gotCourseIds, test_case.expectedOutput.course_ids) {
				t.Errorf("Output: Expected %v but got %v", test_case.expectedOutput,
					gotCourseIds)
			}
		})
	}
}

func TestListSubmittedExercisesByMentee(t *testing.T) {
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
		expectedOutput []models.ListSubmittedExercisesByMentee
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
						AddRow(1, 1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."topic_id" = $1 AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "topic_id"}).
						AddRow(1, 1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "submitted_exercises" WHERE mentee_id = $1 AND "submitted_exercises"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"mentee_id", "exercise_id", "file_name", "file"}).
						AddRow(1, 1, "File Name", "File"))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."id" = $1 AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "question"}).
						AddRow(1, "Question"))
			},
			expectedOutput: []models.ListSubmittedExercisesByMentee{
				{
					ExerciseId: 1,
					FileName:   "File Name",
					File:       "File",
					Question:   "Question",
				},
			},
			expectedError: nil,
		},
		{
			label: "Failed while getting exercise ids",
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
						AddRow(1, 1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."topic_id" = $1 AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnError(errors.New("sql error"))

			},
			expectedOutput: nil,
			expectedError:  errors.New("sql error"),
		},
		{
			label: "Failed while getting list of submitted exercises",
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
						AddRow(1, 1))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "exercises" WHERE "exercises"."topic_id" = $1 AND "exercises"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "topic_id"}).
						AddRow(1, 1))

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
			got, err := DBClient.ListSubmittedExercisesByMentee(test_case.request.mentee_id, test_case.request.course_id)

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
