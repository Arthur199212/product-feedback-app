package feedback_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"product-feedback/feedback"
	mock_feedback "product-feedback/feedback/mocks"
	"product-feedback/validation"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
)

const userIdCtx = "userId"

func Test_CreateFeedback(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userId := 1
		status := "idea"
		createInput := feedback.CreateFeedbackInput{
			Title:    "title",
			Body:     "lorem lorem lorem",
			Category: "ui",
			Status:   &status,
		}

		loggerMock, _ := test.NewNullLogger()
		v := validation.NewValidation()
		serviceMock := mock_feedback.NewMockFeedbackService(ctrl)
		serviceMock.EXPECT().Create(userId, createInput).Return(userId, nil)
		handlerMock := feedback.NewFeedbackHandler(loggerMock, v, serviceMock)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		input := `{"title":"title","body":"lorem lorem lorem","category":"ui","status":"idea"}`
		c.Request, _ = http.NewRequest(
			http.MethodPost,
			"/api/feedback",
			bytes.NewBuffer([]byte(input)),
		)
		c.Set(userIdCtx, userId)

		handlerMock.CreateFeedback(c)

		if w.Code != http.StatusOK {
			t.Fatalf("expected status code %d, but got %d", http.StatusOK, w.Code)
		}
		expectedResponse := fmt.Sprintf(`{"feedbackId":%d}`, userId)
		if w.Body.String() != expectedResponse {
			t.Fatalf("expected response: %v, but got: %v", expectedResponse, w.Body.String())
		}
	})

	t.Run("no userId in context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		loggerMock, _ := test.NewNullLogger()
		v := validation.NewValidation()
		serviceMock := mock_feedback.NewMockFeedbackService(ctrl)
		handlerMock := feedback.NewFeedbackHandler(loggerMock, v, serviceMock)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		input := `{"title":"title","body":"lorem lorem lorem","category":"ui","status":"idea"}`
		c.Request, _ = http.NewRequest(
			http.MethodPost,
			"/api/feedback",
			bytes.NewBuffer([]byte(input)),
		)

		handlerMock.CreateFeedback(c)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("expected status code %d, but got %d", http.StatusOK, w.Code)
		}
		expectedResponse := `{"message":"Unauthorized"}`
		if w.Body.String() != expectedResponse {
			t.Fatalf("expected response: %v, but got: %v", expectedResponse, w.Body.String())
		}
	})

	t.Run("input is invalid JSON", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		loggerMock, _ := test.NewNullLogger()
		v := validation.NewValidation()
		serviceMock := mock_feedback.NewMockFeedbackService(ctrl)
		handlerMock := feedback.NewFeedbackHandler(loggerMock, v, serviceMock)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set(userIdCtx, 1)

		input := `{"title":"title",,"body":"lorem lorem lorem","category":"ui"}`
		c.Request, _ = http.NewRequest(
			http.MethodPost,
			"/api/feedback",
			bytes.NewBuffer([]byte(input)),
		)

		handlerMock.CreateFeedback(c)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected status code %d, but got %d", http.StatusOK, w.Code)
		}
		expectedResponse := `{"message":"Input is invalid"}`
		if w.Body.String() != expectedResponse {
			t.Fatalf("expected response: %v, but got: %v", expectedResponse, w.Body.String())
		}
	})

	t.Run("input is invalid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		loggerMock, _ := test.NewNullLogger()
		v := validation.NewValidation()
		serviceMock := mock_feedback.NewMockFeedbackService(ctrl)
		handlerMock := feedback.NewFeedbackHandler(loggerMock, v, serviceMock)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set(userIdCtx, 1)

		input := `{"title":"no","body":"lorem lorem lorem","category":"ui"}`
		c.Request, _ = http.NewRequest(
			http.MethodPost,
			"/api/feedback",
			bytes.NewBuffer([]byte(input)),
		)

		handlerMock.CreateFeedback(c)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected status code %d, but got %d", http.StatusOK, w.Code)
		}
		expectedResponse := `{"message":"Key: 'CreateFeedbackInput.Title' Error:Field validation for 'Title' failed on the 'min' tag"}`
		if w.Body.String() != expectedResponse {
			t.Fatalf("expected response: %v, but got: %v", expectedResponse, w.Body.String())
		}
	})
}

func Test_DeleteFeedback(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userId := 1

		loggerMock, _ := test.NewNullLogger()
		v := validation.NewValidation()
		serviceMock := mock_feedback.NewMockFeedbackService(ctrl)
		serviceMock.EXPECT().Delete(userId, 1).Return(nil)
		handlerMock := feedback.NewFeedbackHandler(loggerMock, v, serviceMock)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(
			http.MethodDelete,
			"/api/feedback/1",
			nil,
		)
		c.Set(userIdCtx, userId)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

		handlerMock.DeleteFeedback(c)

		if w.Code != http.StatusOK {
			t.Fatalf("expected status code %d, but got %d", http.StatusOK, w.Code)
		}
		expectedResponse := `{"message":"OK"}`
		if w.Body.String() != expectedResponse {
			t.Fatalf("expected response: %v, but got: %v", expectedResponse, w.Body.String())
		}
	})

	t.Run("no userId in context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		loggerMock, _ := test.NewNullLogger()
		v := validation.NewValidation()
		serviceMock := mock_feedback.NewMockFeedbackService(ctrl)
		handlerMock := feedback.NewFeedbackHandler(loggerMock, v, serviceMock)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(
			http.MethodDelete,
			"/api/feedback/1",
			nil,
		)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

		handlerMock.DeleteFeedback(c)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("expected status code %d, but got %d", http.StatusOK, w.Code)
		}
		expectedResponse := `{"message":"Unauthorized"}`
		if w.Body.String() != expectedResponse {
			t.Fatalf("expected response: %v, but got: %v", expectedResponse, w.Body.String())
		}
	})

	t.Run("id param is invalid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userId := 1

		loggerMock, _ := test.NewNullLogger()
		v := validation.NewValidation()
		serviceMock := mock_feedback.NewMockFeedbackService(ctrl)
		handlerMock := feedback.NewFeedbackHandler(loggerMock, v, serviceMock)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(
			http.MethodDelete,
			"/api/feedback/test",
			nil,
		)
		c.Set(userIdCtx, userId)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "test"})

		handlerMock.DeleteFeedback(c)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected status code %d, but got %d", http.StatusOK, w.Code)
		}
		expectedResponse := `{"message":"Invalid feedback id"}`
		if w.Body.String() != expectedResponse {
			t.Fatalf("expected response: %v, but got: %v", expectedResponse, w.Body.String())
		}
	})
}

func Test_GetAllFeedback(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		feedbackList := []feedback.Feedback{
			feedback.Feedback{
				Id:        1,
				Title:     "title",
				Body:      "lorem lorem lorem",
				Category:  "ui",
				Status:    "idea",
				UserId:    1,
				CreatedAt: "2022-08-09 20:29:09.6618642 +0000 UTC",
				UpdatedAt: "2022-08-09 20:29:09.6618642 +0000 UTC",
			},
		}

		loggerMock, _ := test.NewNullLogger()
		v := validation.NewValidation()
		serviceMock := mock_feedback.NewMockFeedbackService(ctrl)
		serviceMock.EXPECT().GetAll().Return(feedbackList, nil)
		handlerMock := feedback.NewFeedbackHandler(loggerMock, v, serviceMock)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(
			http.MethodGet,
			"/api/feedback",
			nil,
		)

		handlerMock.GetAllFeedback(c)

		if w.Code != http.StatusOK {
			t.Fatalf("expected status code %d, but got %d", http.StatusOK, w.Code)
		}
		expectedResponse := `[{"id":1,"title":"title","body":"lorem lorem lorem","category":"ui","status":"idea","userId":1,"createdAt":"2022-08-09 20:29:09.6618642 +0000 UTC","updatedAt":"2022-08-09 20:29:09.6618642 +0000 UTC"}]`
		if w.Body.String() != expectedResponse {
			t.Fatalf("expected response: %v, but got: %v", expectedResponse, w.Body.String())
		}
	})

	t.Run("service returns an error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		loggerMock, _ := test.NewNullLogger()
		v := validation.NewValidation()
		serviceMock := mock_feedback.NewMockFeedbackService(ctrl)
		serviceMock.EXPECT().GetAll().Return(nil, errors.New("test error"))
		handlerMock := feedback.NewFeedbackHandler(loggerMock, v, serviceMock)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(
			http.MethodGet,
			"/api/feedback",
			nil,
		)

		handlerMock.GetAllFeedback(c)

		if w.Code != http.StatusInternalServerError {
			t.Fatalf("expected status code %d, but got %d", http.StatusOK, w.Code)
		}
		expectedResponse := `{"message":"Internal server error"}`
		if w.Body.String() != expectedResponse {
			t.Fatalf("expected response: %v, but got: %v", expectedResponse, w.Body.String())
		}
	})
}

func Test_GetFeedbackById(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		feedbackItem := feedback.Feedback{
			Id:        1,
			Title:     "title",
			Body:      "lorem lorem lorem",
			Category:  "ui",
			Status:    "idea",
			UserId:    1,
			CreatedAt: "2022-08-09 20:29:09.6618642 +0000 UTC",
			UpdatedAt: "2022-08-09 20:29:09.6618642 +0000 UTC",
		}

		loggerMock, _ := test.NewNullLogger()
		v := validation.NewValidation()
		serviceMock := mock_feedback.NewMockFeedbackService(ctrl)
		serviceMock.EXPECT().GetById(1).Return(feedbackItem, nil)
		handlerMock := feedback.NewFeedbackHandler(loggerMock, v, serviceMock)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(
			http.MethodGet,
			"/api/feedback/1",
			nil,
		)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

		handlerMock.GetFeedbackById(c)

		if w.Code != http.StatusOK {
			t.Fatalf("expected status code %d, but got %d", http.StatusOK, w.Code)
		}
		expectedResponse := `{"id":1,"title":"title","body":"lorem lorem lorem","category":"ui","status":"idea","userId":1,"createdAt":"2022-08-09 20:29:09.6618642 +0000 UTC","updatedAt":"2022-08-09 20:29:09.6618642 +0000 UTC"}`
		if w.Body.String() != expectedResponse {
			t.Fatalf("expected response: %v, but got: %v", expectedResponse, w.Body.String())
		}
	})

	t.Run("id param is invalid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		loggerMock, _ := test.NewNullLogger()
		v := validation.NewValidation()
		serviceMock := mock_feedback.NewMockFeedbackService(ctrl)
		handlerMock := feedback.NewFeedbackHandler(loggerMock, v, serviceMock)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(
			http.MethodGet,
			"/api/feedback/test",
			nil,
		)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "test"})

		handlerMock.GetFeedbackById(c)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected status code %d, but got %d", http.StatusBadRequest, w.Code)
		}
		expectedResponse := `{"message":"Invalid feedback id"}`
		if w.Body.String() != expectedResponse {
			t.Fatalf("expected response: %v, but got: %v", expectedResponse, w.Body.String())
		}
	})
}