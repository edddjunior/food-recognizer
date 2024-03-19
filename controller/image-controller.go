package controller

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"

	"github.com/edddjunior/food-recognizer/service"
)

type ImageController struct {
	service *service.ImageService
	*gin.Context
}

func NewImageController(service *service.ImageService) *ImageController {
	return &ImageController{
		service: service,
	}
}

func (c *ImageController) InitRoutes() {
	app := gin.Default()
	api := app.Group("/api/images")

	api.POST("/process", c.processImage)

	app.Run(":3000")
}

func (c *ImageController) processImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	tempFile, err := os.CreateTemp("", "uploaded-image-*.jpg")
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("create temp file err: %s", err.Error()))
		return
	}
	defer os.Remove(tempFile.Name())

	if err := ctx.SaveUploadedFile(file, tempFile.Name()); err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	cmd := exec.Command("python3", "../model/runner.py", tempFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("process image err: %s", err.Error()))
		return
	}

	outputString := string(output)

	ctx.String(http.StatusOK, outputString)
}
