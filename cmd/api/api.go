package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nuriansyah/api-gateway/cmd/config"
	"github.com/nuriansyah/api-gateway/internal/repository"
)

type API struct {
	commentRepo          repository.CommentsRepository
	userRepo             repository.UserRepository
	postRepo             repository.PostRepository
	mentorshipRepository repository.MentorshipRepository
	router               *gin.Engine
}

func NewAPi(userRepo repository.UserRepository, postRepo repository.PostRepository, mentorshipRepository repository.MentorshipRepository) API {
	router := gin.Default()
	api := &API{
		userRepo:             userRepo,
		postRepo:             postRepo,
		mentorshipRepository: mentorshipRepository,
		router:               router,
	}

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.POST("/login", api.login)
	router.POST("/register", api.register)
	router.GET("/getMentored/:id", api.readMentored)
	router.GET("/getMentored", api.readMentoreds)

	router.GET("/api/post/:id", api.readPosts)
	postRouter := router.Group("/api/post", AuthMiddleware())
	{
		postRouter.POST("/create", api.createPost)
	}

	router.Use(gin.Recovery())

	return *api

}

func (api *API) Handler() *gin.Engine {
	return api.router
}

func (api *API) Start() {
	setPort := config.New(".env")
	api.Handler().Run(setPort.Get("APP_PORT"))
}
