package feedback_test

import (
	"database/sql"
	"errors"
	"product-feedback/feedback"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRepository_Create(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var expectedError error = nil
		expectedFeedbackId := 1
		userId := 1
		input := feedback.CreateFeedbackInput{
			Title:    "title",
			Body:     "lorem lorem lorem",
			Category: "ui",
			Status:   nil,
		}

		rowsMock := sqlmock.NewRows([]string{"id"}).AddRow(expectedFeedbackId)
		mock.ExpectQuery("INSERT INTO feedback").WithArgs(
			input.Title,
			input.Body,
			input.Category,
			"idea",
			userId,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).WillReturnRows(rowsMock)

		repo := feedback.NewFeedbackRepository(db)

		feedbackId, err := repo.Create(userId, input)
		if err != expectedError {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
		if feedbackId != expectedFeedbackId {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
	})

	t.Run("input has Status field", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var expectedError error = nil
		expectedFeedbackId := 1
		userId := 1
		status := "defined"
		input := feedback.CreateFeedbackInput{
			Title:    "title",
			Body:     "lorem lorem lorem",
			Category: "ui",
			Status:   &status,
		}

		rowsMock := sqlmock.NewRows([]string{"id"}).AddRow(expectedFeedbackId)
		mock.ExpectQuery("INSERT INTO feedback").WithArgs(
			input.Title,
			input.Body,
			input.Category,
			status,
			userId,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).WillReturnRows(rowsMock)

		repo := feedback.NewFeedbackRepository(db)

		feedbackId, err := repo.Create(userId, input)
		if err != expectedError {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
		if feedbackId != expectedFeedbackId {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
	})

	t.Run("returns an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var expectedError error = errors.New("test error")
		expectedFeedbackId := 0
		userId := 1
		status := "defined"
		input := feedback.CreateFeedbackInput{
			Title:    "title",
			Body:     "lorem lorem lorem",
			Category: "ui",
			Status:   &status,
		}

		mock.ExpectQuery("INSERT INTO feedback").WithArgs(
			input.Title,
			input.Body,
			input.Category,
			status,
			userId,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).WillReturnError(expectedError)

		repo := feedback.NewFeedbackRepository(db)

		feedbackId, err := repo.Create(userId, input)
		if err != expectedError {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
		if feedbackId != expectedFeedbackId {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
	})
}

func TestRepository_Delete(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var expectedError error = nil
		userId := 1
		feedbackId := 1

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM comments").WithArgs(
			feedbackId,
		).WillReturnResult(sqlmock.NewResult(int64(1), 1))
		mock.ExpectExec("DELETE FROM votes").WithArgs(
			feedbackId,
		).WillReturnResult(sqlmock.NewResult(int64(1), 1))
		mock.ExpectExec("DELETE FROM feedback").WithArgs(
			userId,
			feedbackId,
		).WillReturnResult(sqlmock.NewResult(int64(feedbackId), 1))
		mock.ExpectCommit()

		repo := feedback.NewFeedbackRepository(db)

		err = repo.Delete(userId, feedbackId)
		if err != expectedError {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
	})

	t.Run("comments delete returns an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		expectedError := errors.New("test error")
		userId := 1
		feedbackId := 1

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM comments").WithArgs(
			feedbackId,
		).WillReturnError(expectedError)
		mock.ExpectRollback()

		repo := feedback.NewFeedbackRepository(db)

		err = repo.Delete(userId, feedbackId)
		if err != expectedError {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
	})

	t.Run("votes delete returns an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		expectedError := errors.New("test error")
		userId := 1
		feedbackId := 1

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM comments").WithArgs(
			feedbackId,
		).WillReturnResult(sqlmock.NewResult(int64(1), 1))
		mock.ExpectExec("DELETE FROM votes").WithArgs(
			feedbackId,
		).WillReturnError(expectedError)
		mock.ExpectRollback()

		repo := feedback.NewFeedbackRepository(db)

		err = repo.Delete(userId, feedbackId)
		if err != expectedError {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
	})

	t.Run("feedback delete returns an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		expectedError := errors.New("test error")
		userId := 1
		feedbackId := 1

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM comments").WithArgs(
			feedbackId,
		).WillReturnResult(sqlmock.NewResult(int64(1), 1))
		mock.ExpectExec("DELETE FROM votes").WithArgs(
			feedbackId,
		).WillReturnResult(sqlmock.NewResult(int64(1), 1))
		mock.ExpectExec("DELETE FROM feedback").WithArgs(
			userId,
			feedbackId,
		).WillReturnError(expectedError)
		mock.ExpectRollback()

		repo := feedback.NewFeedbackRepository(db)

		err = repo.Delete(userId, feedbackId)
		if err != expectedError {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
	})

	t.Run("no rows in result set", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		expectedError := sql.ErrNoRows
		userId := 1
		feedbackId := 1

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM comments").WithArgs(
			feedbackId,
		).WillReturnResult(sqlmock.NewResult(int64(1), 1))
		mock.ExpectExec("DELETE FROM votes").WithArgs(
			feedbackId,
		).WillReturnResult(sqlmock.NewResult(int64(1), 1))
		mock.ExpectExec("DELETE FROM feedback").WithArgs(
			userId,
			feedbackId,
		).WillReturnResult(sqlmock.NewResult(int64(feedbackId), 0))
		mock.ExpectRollback()

		repo := feedback.NewFeedbackRepository(db)

		err = repo.Delete(userId, feedbackId)
		if err != expectedError {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
	})
}

func TestRepository_GetAll(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		repo := feedback.NewFeedbackRepository(db)

		var expectedError error = nil
		timeString := time.Now().Add(-time.Hour * 20).UTC().String()
		title1 := "title #1"
		title2 := "title #2"
		rows := sqlmock.NewRows(
			[]string{"id", "title", "body", "category", "status", "user_id", "created_at", "updated_at"},
		).AddRow(
			1, title1, "lorem lorem lorem", "ui", "idea", 1, timeString, timeString,
		).AddRow(
			1, title2, "lorem lorem lorem", "ui", "defined", 1, timeString, timeString,
		)
		mock.ExpectQuery("SELECT .* FROM feedback").WithArgs().WillReturnRows(rows)

		fList, err := repo.GetAll()
		if err != expectedError {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
		if fList[0].Title != title1 && fList[1].Title != title2 {
			t.Fatalf("returned items do not match data from DB")
		}
	})
}

func TestRepository_GetById(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		repo := feedback.NewFeedbackRepository(db)

		var expectedError error = nil
		feedbackId := 1
		timeString := time.Now().Add(-time.Hour * 20).UTC().String()
		rows := sqlmock.NewRows(
			[]string{"id", "title", "body", "category", "status", "user_id", "created_at", "updated_at"},
		).AddRow(1, "title", "lorem lorem lorem", "ui", "idea", 1, timeString, timeString)
		mock.ExpectQuery("SELECT .* FROM feedback").WithArgs(
			feedbackId,
		).WillReturnRows(rows)

		f, err := repo.GetById(feedbackId)
		if err != expectedError {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
		if f.Id != feedbackId {
			t.Fatalf("expected to get feedback with id %d, but got %d", feedbackId, f.Id)
		}
	})
}

func TestRepository_Update(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		repo := feedback.NewFeedbackRepository(db)

		var expectedError error = nil
		userId := 1
		feedbackId := 1
		title := "title"
		body := "lorem lorem lorem"
		category := "ux"
		status := "done"
		input := feedback.UpdateFeedbackInput{
			Title:    &title,
			Body:     &body,
			Category: &category,
			Status:   &status,
		}

		mock.ExpectExec("UPDATE feedback").WithArgs(
			body,
			category,
			status,
			title,
			sqlmock.AnyArg(),
			userId,
			feedbackId,
		).WillReturnResult(sqlmock.NewResult(int64(1), 1))

		err = repo.Update(userId, feedbackId, input)
		if err != expectedError {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
	})

	t.Run("no input to update", func(t *testing.T) {
		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		repo := feedback.NewFeedbackRepository(db)

		expectedError := feedback.ErrNoInputToUpdate
		userId := 1
		feedbackId := 1
		input := feedback.UpdateFeedbackInput{}

		err = repo.Update(userId, feedbackId, input)
		if err != expectedError {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
	})

	t.Run("no rows in result set", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		repo := feedback.NewFeedbackRepository(db)

		expectedError := sql.ErrNoRows
		userId := 1
		feedbackId := 1
		title := "title"
		body := "lorem lorem lorem"
		category := "ux"
		status := "done"
		input := feedback.UpdateFeedbackInput{
			Title:    &title,
			Body:     &body,
			Category: &category,
			Status:   &status,
		}

		mock.ExpectExec("UPDATE feedback").WithArgs(
			body,
			category,
			status,
			title,
			sqlmock.AnyArg(),
			userId,
			feedbackId,
		).WillReturnResult(sqlmock.NewResult(int64(1), 0))

		err = repo.Update(userId, feedbackId, input)
		if err != expectedError {
			t.Fatalf("expected error %v, but got %v", expectedError, err)
		}
	})
}
