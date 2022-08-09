package feedback_test

import (
	"bytes"
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
