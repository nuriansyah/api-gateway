package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nuriansyah/api-gateway/cmd/api"
	"github.com/nuriansyah/api-gateway/cmd/config"
	"github.com/nuriansyah/api-gateway/internal/repository"
	"net/http"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
func main() {
	configuration := config.New(".env")
	db, err := config.NewInitializedDatabase(configuration)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)
	mentorshipRepo := repository.NewMentorshipRepository(db)

	mainAPI := api.NewAPi(*userRepo, *postRepo, *mentorshipRepo)
	mainAPI.Start()
}
