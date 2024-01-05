package database

import (
	"errors"
	"lpt/pkg/models"
	reflect "reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddUser(t *testing.T) {
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}
	test_cases := []struct {
		label          string
		request        models.User
		mockFunc       func()
		expectedOutput int32
		expectedError  error
	}{
		{
			label: "Success",
			request: models.User{
				Name:  "User Name",
				Email: "User Email",
				Role:  0,
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("created_at","updated_at","deleted_at","name","email","role") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "User Name", "User Email", 0).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
			expectedOutput: 1,
			expectedError:  nil,
		},
		{
			label: "Failure",
			request: models.User{
				Name:  "User Name",
				Email: "User Email",
				Role:  0,
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("created_at","updated_at","deleted_at","name","email","role") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "User Name", "User Email", 0).
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
			got, err := DBClient.AddUser(test_case.request)

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
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}
	test_cases := []struct {
		label          string
		request        int32
		mockFunc       func()
		expectedOutput models.User
		expectedError  error
	}{
		{
			label:   "Success",
			request: 1,
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"name", "email", "role"}).
						AddRow("User Name", "User Email", 1))
			},
			expectedOutput: models.User{
				Name:  "User Name",
				Email: "User Email",
				Role:  1,
			},
			expectedError: nil,
		},
		{
			label:   "Failure",
			request: 1,
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnError(errors.New("There is no user"))

			},
			expectedOutput: models.User{},
			expectedError:  errors.New("There is no user"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.GetUserDetails(test_case.request)

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
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}
	test_cases := []struct {
		label          string
		request        int32
		mockFunc       func()
		expectedOutput []string
		expectedError  error
	}{
		{
			label:   "Success",
			request: 1,
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT users.name FROM "users" WHERE users.role = $1`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow("User1").AddRow("User2"))
			},
			expectedOutput: []string{"User1", "User2"},
			expectedError:  nil,
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.ListUsersByRole(test_case.request)

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
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}
	test_cases := []struct {
		label          string
		request        models.CoursesAssigned
		mockFunc       func()
		expectedOutput int32
		expectedError  error
	}{
		{
			label: "Success",
			request: models.CoursesAssigned{
				MentorId: 1,
				MenteeId: 2,
				CourseId: 3,
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "courses_assigned" ("created_at","updated_at","deleted_at","mentor_id","mentee_id","course_id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1, 2, 3).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
			expectedOutput: 1,
			expectedError:  nil,
		},
		{
			label: "Failure",
			request: models.CoursesAssigned{
				MentorId: 1,
				MenteeId: 2,
				CourseId: 3,
			},
			mockFunc: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "courses_assigned" ("created_at","updated_at","deleted_at","mentor_id","mentee_id","course_id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1, 2, 3).
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
			got, err := DBClient.CreateAssignment(test_case.request)

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
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}

	coursesAssigned := models.CoursesAssigned{
		MentorId: 1,
		MenteeId: 2,
		CourseId: 3,
	}

	coursesAssigned.ID = 1
	coursesAssigned.Mentor.ID = 1
	coursesAssigned.Mentee.ID = 2
	coursesAssigned.Course.ID = 3

	test_cases := []struct {
		label          string
		mockFunc       func()
		expectedOutput []models.CoursesAssigned
		expectedError  error
	}{
		{
			label: "Success",
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses_assigned" WHERE "courses_assigned"."deleted_at" IS NULL`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "mentor_id", "mentee_id", "course_id"}).
						AddRow(1, 1, 2, 3))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses" WHERE "courses"."id" = $1 AND "courses"."deleted_at" IS NULL`)).
					WithArgs(3).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(3))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`)).
					WithArgs(2).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(2))

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))
			},
			expectedOutput: []models.CoursesAssigned{
				coursesAssigned,
			},
			expectedError: nil,
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.ListCurrentAssignments()

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
func TestGetUserIdByEmail(t *testing.T) {
	db, mock := MockDB()
	defer mock.ExpectClose()

	DBClient := DBClient{
		DB: db,
	}
	test_cases := []struct {
		label          string
		request        string
		mockFunc       func()
		expectedOutput int32
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
			},
			expectedOutput: 1,
			expectedError:  nil,
		},
		{
			label:   "Failure",
			request: "User Email",
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM "users" WHERE email = $1`)).
					WithArgs("User Email").
					WillReturnError(errors.New("There is no user"))

			},
			expectedOutput: 0,
			expectedError:  errors.New("There is no user"),
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.label, func(t *testing.T) {
			test_case.mockFunc()
			got, err := DBClient.GetUserIdByEmail(test_case.request)

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
