package core

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/AndreyDyubin/app/storage"
	"bytes"
	"github.com/satori/go.uuid"
	"strings"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"io"
	"gopkg.in/reform.v1"
	"go.uber.org/zap"
)

var UploadService = new(uploadService)

type uploadService struct {
	logger *zap.SugaredLogger
	db     *reform.DB
}

func NewUploadService(db *reform.DB, logger *zap.SugaredLogger) *uploadService {
	UploadService.db = db
	UploadService.logger = logger
	return UploadService
}

type result struct {
	FileID string `json:"file_id"`
}
type resultFile struct {
	Name string
	Type string
	Body io.Reader
}

func (us *uploadService) Upload(name string, data []byte) (*result, error) {

	u := uuid.NewV4().String()
	u = strings.Replace(u, "-", "", -1)
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)
	path := bytes.Buffer{}
	path.WriteString("/media/")
	path.WriteString(u[:3])
	path.WriteString("/")
	path.WriteString(u[3:])
	params := &s3.PutObjectInput{
		Bucket:        aws.String("maddevilbucket"),
		Key:           aws.String(path.String()),
		Body:          fileBytes,
		ContentLength: aws.Int64(int64(len(data))),
		ContentType:   aws.String(fileType),
	}
	_, err := storage.S3.PutObject(params)
	if err != nil {
		us.logger.Error(err)
		return nil, err
	}
	ID, err := storage.SaveDataFile(u, name)

	return &result{ID}, err
}

func (us *uploadService) Download(ID string) (*resultFile, error) {
	df, err := storage.DataFile(ID)
	if err != nil {
		us.logger.Error(err)
		return nil, err
	}
	path := bytes.Buffer{}
	path.WriteString("/media/")
	path.WriteString(df.UUID[:3])
	path.WriteString("/")
	path.WriteString(df.UUID[3:])
	key := path.String()
	params := &s3.GetObjectInput{
		Bucket: aws.String("maddevilbucket"),
		Key:    &key,
	}

	resp, err := storage.S3.GetObject(params)
	if err != nil {
		us.logger.Error(err)
		return nil, err
	}
	return &resultFile{df.Name, *resp.ContentType, resp.Body}, err
}
