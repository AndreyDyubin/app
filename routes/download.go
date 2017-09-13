package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/AndreyDyubin/app/storage"
	"strconv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"bytes"
)

func Download(c echo.Context) error {
	IDs := c.QueryParam("id")
	ID, err := strconv.ParseInt(IDs, 10, 64)
	if err != nil {
		return c.HTML(http.StatusOK, "Неверынй ID, "+IDs)
	}
	df, err := storage.DataFile(ID)
	if err != nil {
		return c.HTML(http.StatusOK, err.Error())
	}
	key := "/media/" + df.Name
	params := &s3.GetObjectInput{
		Bucket: aws.String("testBucket"),
		Key:    &key,
	}

	resp, err := storage.S3.GetObject(params)
	if err != nil {
		return c.HTML(http.StatusOK, err.Error())
	}
	buf := bytes.Buffer{}
	buf.WriteString("attachment; filename=\"")
	buf.WriteString(df.Name)
	buf.WriteString("\"")
	c.Response().Header().Add("Content-Disposition", buf.String())
	return c.Stream(http.StatusOK, *resp.ContentType, resp.Body)
}
