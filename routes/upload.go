package routes

import (
	"github.com/labstack/echo"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"bytes"
	"net/http"
	"github.com/AndreyDyubin/app/storage"
	"strconv"
)

type result struct {
	FileID string `json:"file_id"`
}

func Upload(c echo.Context) error {
	var err error
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	size := file.Size
	buffer := make([]byte, size) // read file content to buffer

	src.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := "/media/" + file.Filename
	params := &s3.PutObjectInput{
		Bucket:        aws.String("testBucket"),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	_, err = storage.S3.PutObject(params)
	if err != nil {
		return c.JSON(http.StatusOK, &result{})
	}
	ID, err := storage.SaveDataFile(file.Filename)
	if err != nil {
		return c.JSON(http.StatusOK, &result{})
	}
	res := result{strconv.FormatInt(ID, 10)}
	return c.JSON(http.StatusOK, &res)
}
