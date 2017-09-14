package routes

import (
	"github.com/labstack/echo"
	"github.com/AndreyDyubin/app/storage"
	"net/http"
)

func List(c echo.Context) error {
	fileList, err := storage.SelectList()
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, &fileList)
}
