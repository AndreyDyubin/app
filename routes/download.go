package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"bytes"
	"github.com/AndreyDyubin/app/core"
)

func Download(c echo.Context) error {
	ID := c.QueryParam("id")

	rf, err := core.Download(ID)
	if err != nil {
		return c.HTML(http.StatusOK, err.Error())
	}
	buf := bytes.Buffer{}
	buf.WriteString("attachment; filename=\"")
	buf.WriteString(rf.Name)
	buf.WriteString("\"")
	c.Response().Header().Add("Content-Disposition", buf.String())
	return c.Stream(http.StatusOK, rf.Type, rf.Body)
}
