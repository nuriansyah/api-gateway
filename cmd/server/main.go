package main

import (
	"github.com/nuriansyah/api-gateway/cmd/api"
	"github.com/nuriansyah/api-gateway/cmd/config"
	"github.com/nuriansyah/api-gateway/internal/repository"
)

func main() {
	configuration := config.New("../../.env")
	db, err := config.NewInitializedDatabase(configuration)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)

	mainAPI := api.NewAPi(*userRepo, *postRepo)
	mainAPI.Start()
}
