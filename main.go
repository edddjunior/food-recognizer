package main

import (
	"github.com/edddjunior/food-recognizer/controller"
	"github.com/edddjunior/food-recognizer/service"
)

func main() {
	imageService := service.NewImageService()
	imageController := controller.NewImageController(imageService)
	imageController.InitRoutes()
}
