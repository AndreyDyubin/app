package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/AndreyDyubin/app/core"
)

func Upload(c echo.Context) error {
	var err error
	file, err := c.FormFile("file")
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	src, err := file.Open()
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	defer src.Close()
	size := file.Size
	buffer := make([]byte, size)

	src.Read(buffer)

	res, err := core.UploadService.Upload(file.Filename, buffer)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, &res)
}
