package comment

import (
	"database/sql"
	"fmt"
	"net/http"
	"product-feedback/middleware"
	"product-feedback/validation"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CommentHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type commentHandler struct {
	l       *logrus.Logger
	v       *validation.Validation
	service CommentService
}

func NewCommentHandler(
	l *logrus.Logger,
	v *validation.Validation,
	service CommentService,
) CommentHandler {
	return &commentHandler{
		l:       l,
		v:       v,
		service: service,
	}
}

// swagger:route GET /api/comments comments getAllComments
// Returns a list of all comments in the system
//
// security:
// - Bearer:
//
// responses:
//	200: getAllCommentsResponse

func (h *commentHandler) getAllComments(c *gin.Context) {
	// todo: implement options:
	// filter by: userId
	// sorted: date of creation, date of update
	// pagination: limit/size=<uint>, page=<uint>

	feedbackIds, err := parseFeedbackIdsFromQuery(c.Query("feedbackId"))
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	comments, err := h.service.GetAll(feedbackIds)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func parseFeedbackIdsFromQuery(str string) ([]int, error) {
	if str == "" {
		return []int{}, nil
	}

	ids := strings.Split(str, ",")
	parsedIds := make([]int, len(ids))
	for i := range ids {
		parsedId, err := strconv.Atoi(ids[i])
		if err != nil {
			return parsedIds, fmt.Errorf("invalid feedbackId param")
		}
		parsedIds[i] = parsedId
	}
	return parsedIds, nil
}

type createCommentInput struct {
	// Body of the comment
	//
	// required: true
	// min length: 5
	// max length: 255
	Body string `json:"body" validate:"required,min=5,max=255"`
	// Id of the feedback this comment is related to
	//
	// required: true
	// min: 1
	FeedbackId int `json:"feedbackId" validate:"required,gt=0"`
	// Id of the comment that this comment relates to
	//
	// required: false
	// min: 1
	ParentId *int `json:"parentId" db:"parent_id" validate:"omitempty,gt=0"`
}

// swagger:route POST /api/comments comments createComment
// Creates a comment
//
// security:
// - Bearer:
//
// responses:
//	200: createCommentResponse

func (h *commentHandler) createComment(c *gin.Context) {
	userId, err := middleware.GetUserIdFromGinCtx(c)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
		return
	}

	var input createCommentInput
	if err = c.BindJSON(&input); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Input is invalid",
		})
		return
	}

	if err = h.v.ValidateStruct(input); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	commentId, err := h.service.Create(userId, input)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"commentId": commentId,
	})
}

func (h *commentHandler) deleteComment(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "deleteComment not implemented",
	})
}

// swagger:route GET /api/comments/:id comments getCommentById
// Returns comment by id
//
// security:
// - Bearer:
//
// responses:
//	200: getCommentByIdResponse
//	404: errorResponse

func (h *commentHandler) getCommentById(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "commentId is invalid",
		})
		return
	}

	if err = h.v.ValidateVar(commentId, "required,gt=0"); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "commentId is invalid",
		})
		return
	}

	comment, err := h.service.GetById(commentId)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
			"message": "Comment not found",
		})
		return
	default:
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *commentHandler) updateComment(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "updateComment not implemented",
	})
}
