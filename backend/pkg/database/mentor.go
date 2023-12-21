package database

import (
	"lpt/pkg/models"
	"slices"
)

func (db DBClient) ListAssignedMenteesAndCourses(mentor_id int32) ([]string, []int32, []int32, error) {
	assigned_details := []models.CoursesAssigned{}
	res := db.DB.
		Where("mentor_id = ?", mentor_id).
		Find(&assigned_details)

	if res.RowsAffected == 0 || res.Error != nil {
		return nil, nil, nil, res.Error
	}

	menteeEmails := []string{}
	menteeIds := []int32{}
	courseIds := []int32{}

	for _, assigned_detail := range assigned_details {
		var menteeEmail string
		res := db.DB.
			Table("users").
			Select("email").
			Where("id = ?", assigned_detail.MenteeId).
			Find(&menteeEmail)

		if res.RowsAffected == 0 || res.Error != nil {
			return nil, nil, nil, res.Error
		}

		menteeEmails = append(menteeEmails, menteeEmail)
		menteeIds = append(menteeIds, assigned_detail.MenteeId)
		courseIds = append(courseIds, assigned_detail.CourseId)
	}

	return menteeEmails, menteeIds, courseIds, nil
}

func (db DBClient) ListSubmittedExercisesByMentee(mentee_id int32, course_id int32) ([]models.ListSubmittedExercisesByMentee, error) {
	exercise_Ids, err := db.ListExerciseIds(course_id)
	if err != nil {
		return nil, err
	}

	submittedExercises, err := db.ListSubmittedExercises(mentee_id)
	if err != nil {
		return nil, err
	}

	var submittedExerciseDetails []models.ListSubmittedExercisesByMentee

	for _, submittedExercise := range submittedExercises {
		if slices.Contains(exercise_Ids, submittedExercise.ExerciseId) {
			submittedExerciseDetails = append(submittedExerciseDetails,
				models.ListSubmittedExercisesByMentee{
					ExerciseId: submittedExercise.ExerciseId,
					FileName:   submittedExercise.FileName,
					File:       submittedExercise.File,
					Question:   submittedExercise.Exercise.Question,
				})
		}
	}

	return submittedExerciseDetails, nil
}
