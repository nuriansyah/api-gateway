package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestMentorship struct {
	MahasiswaID int `json:"mahasiswaID"`
	DosenID     int `json:"dosenID"`
}

type Response struct {
	Message string `json:"message"`
}

type ResponseMentored struct {
	ID      int     `json:"ID"`
	Name    string  `json:"name"`
	Company *string `json:"company"`
	Program string  `json:"program"`
	Batch   *int    `json:"batch"`
	Nrp     string  `json:"nrp"`
	Prodi   string  `json:"prodi"`
}

func (api API) readMentored(c *gin.Context) {
	userId, err := api.getUserIdFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, Response{"Unauthorized"})
		return
	}

	mentoredList, err := api.mentorshipRepository.GetUserDataMentorship(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Message: err.Error()})
		return
	}

	var mentoredListResponse []ResponseMentored
	for i, mentorship := range mentoredList {
		mentoredListResponse = append(mentoredListResponse, ResponseMentored{
			ID:      i + 1,
			Name:    mentorship.Name,
			Company: mentorship.Company,
			Program: mentorship.Program,
			Batch:   mentorship.Batch,
			Nrp:     mentorship.Nrp,
			Prodi:   mentorship.Prodi,
		})
	}

	c.JSON(http.StatusOK, mentoredListResponse)
}

func (api API) readMentoreds(c *gin.Context) {
	mentoredList, err := api.mentorshipRepository.GetUserDataMentorships()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Message: err.Error()})
		return
	}

	var mentoredListResponse []ResponseMentored
	for _, mentorship := range mentoredList {
		mentoredListResponse = append(mentoredListResponse, ResponseMentored{
			ID:      mentorship.Id,
			Name:    mentorship.Name,
			Company: mentorship.Company,
		})
	}

	c.JSON(http.StatusOK, mentoredListResponse)
}
