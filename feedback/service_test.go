package feedback_test

import (
	"database/sql"
	"errors"
	"product-feedback/feedback"
	mock_feedback "product-feedback/feedback/mocks"
	"product-feedback/notifier"
	"time"

	mock_notifier "product-feedback/notifier/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

type NotifierMock struct {
	notifier.NotifierService
}

func TestService_Create(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedId := 1
		var expectedError error = nil
		input := feedback.CreateFeedbackInput{
			Title:    "title",
			Body:     "lorem lorem lorem",
			Category: "defect",
			Status:   nil,
		}
		notifierServiceMock := mock_notifier.NewMockNotifierService(ctrl)
		notifierServiceMock.EXPECT().BroadcastMessage(notifier.CreateEvent, notifier.SubjectFeedback, 1).MaxTimes(1)
		repo := mock_feedback.NewMockFeedbackRepository(ctrl)
		repo.EXPECT().Create(1, input).Return(expectedId, expectedError)
		service := feedback.NewFeedbackService(repo, notifierServiceMock)

		id, err := service.Create(1, input)

		// let gorutine with BroadcastMessage to complete
		time.Sleep(time.Millisecond)

		if err != expectedError {
			t.Fatalf("expected err to be %v, but got %v", expectedError, err)
		}
		if id != expectedId {
			t.Fatalf("expected to get id %d, but got %d", expectedId, id)
		}
	})

	t.Run("returns an error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedId := 0
		var expectedError error = errors.New("test error")
		input := feedback.CreateFeedbackInput{
			Title:    "title",
			Body:     "lorem lorem lorem",
			Category: "defect",
			Status:   nil,
		}
		notifierServiceMock := mock_notifier.NewMockNotifierService(ctrl)
		repo := mock_feedback.NewMockFeedbackRepository(ctrl)
		repo.EXPECT().Create(1, input).Return(expectedId, expectedError)

		service := feedback.NewFeedbackService(repo, notifierServiceMock)

		id, err := service.Create(1, input)
		if err != expectedError {
			t.Fatalf("expected err to be %v, but got %v", expectedError, err)
		}
		if id != expectedId {
			t.Fatalf("expected to get id %d, but got %d", expectedId, id)
		}
	})
}

func TestService_Delete(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var expectedError error = nil
		notifierServiceMock := mock_notifier.NewMockNotifierService(ctrl)
		notifierServiceMock.EXPECT().BroadcastMessage(notifier.DeleteEvent, notifier.SubjectFeedback, 1).MaxTimes(1)
		repo := mock_feedback.NewMockFeedbackRepository(ctrl)
		repo.EXPECT().Delete(1, 1).Return(expectedError)
		repo.EXPECT().GetById(1).Return(feedback.Feedback{}, nil)
		service := feedback.NewFeedbackService(repo, notifierServiceMock)

		err := service.Delete(1, 1)

		// let gorutine with BroadcastMessage to complete
		time.Sleep(time.Millisecond)

		if err != expectedError {
			t.Fatalf("expected error to be %v, but got %v", expectedError, err)
		}
	})

	t.Run("feedback is not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var expectedError error = sql.ErrNoRows
		notifierServiceMock := mock_notifier.NewMockNotifierService(ctrl)
		repo := mock_feedback.NewMockFeedbackRepository(ctrl)
		repo.EXPECT().GetById(1).Return(feedback.Feedback{}, expectedError)
		service := feedback.NewFeedbackService(repo, notifierServiceMock)

		err := service.Delete(1, 1)
		if err != expectedError {
			t.Fatalf("expected error to be %v, but got %v", expectedError, err)
		}
	})

	t.Run("returns an error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var expectedError error = errors.New("test error")
		notifierServiceMock := mock_notifier.NewMockNotifierService(ctrl)
		repo := mock_feedback.NewMockFeedbackRepository(ctrl)
		repo.EXPECT().GetById(1).Return(feedback.Feedback{}, nil)
		repo.EXPECT().Delete(1, 1).Return(expectedError)
		service := feedback.NewFeedbackService(repo, notifierServiceMock)

		err := service.Delete(1, 1)
		if err != expectedError {
			t.Fatalf("expected error to be %v, but got %v", expectedError, err)
		}
	})
}

func TestService_Update(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var expectedError error = nil
		title := "title"
		input := feedback.UpdateFeedbackInput{
			Title: &title,
		}
		notifierServiceMock := mock_notifier.NewMockNotifierService(ctrl)
		notifierServiceMock.EXPECT().BroadcastMessage(notifier.UpdateEvent, notifier.SubjectFeedback, 1).MaxTimes(1)
		repo := mock_feedback.NewMockFeedbackRepository(ctrl)
		repo.EXPECT().Update(1, 1, input).Return(expectedError)
		service := feedback.NewFeedbackService(repo, notifierServiceMock)

		err := service.Update(1, 1, input)

		// let gorutine with BroadcastMessage to complete
		time.Sleep(time.Millisecond)

		if err != expectedError {
			t.Fatalf("expected error to be %v, but got %v", expectedError, err)
		}
	})

	t.Run("returns an error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var expectedError error = errors.New("test error")
		title := "title"
		input := feedback.UpdateFeedbackInput{
			Title: &title,
		}
		notifierServiceMock := mock_notifier.NewMockNotifierService(ctrl)
		repo := mock_feedback.NewMockFeedbackRepository(ctrl)
		repo.EXPECT().Update(1, 1, input).Return(expectedError)
		service := feedback.NewFeedbackService(repo, notifierServiceMock)

		err := service.Update(1, 1, input)
		if err != expectedError {
			t.Fatalf("expected error to be %v, but got %v", expectedError, err)
		}
	})
}
