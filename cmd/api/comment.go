package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nuriansyah/api-gateway/internal/repository"
	"github.com/nuriansyah/api-gateway/utils"
	"net/http"
	"strconv"
)

type CreateCommentRequest struct {
	PostID  int    `json:"post_id" binding:"required,number"`
	Comment string `json:"comment" binding:"required"`
}

func (api API) CreateComment(c *gin.Context) {
	var createCommentRequest CreateCommentRequest
	err := c.ShouldBind(&createCommentRequest)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"errors": utils.GetErrorMessage(ve)},
			)
		} else {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)
		}
		return
	}
	userID, err := api.getUserIdFromToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = api.commentRepo.InsertComment(repository.Comment{
		PostID:   createCommentRequest.PostID,
		AuthorID: userID,
		Comment:  createCommentRequest.Comment,
	})

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"message": "Add Comment Successful"},
	)
}
func (api *API) ReadAllComment(c *gin.Context) {
	postID, err := strconv.Atoi(c.Query("postID"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	userID := -1
	if c.GetHeader("Authorization") != "" {
		userID, err = api.getUserIdFromToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	comments, err := api.commentRepo.SelectAllCommentsByID(userID, postID)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		comments,
	)
}
