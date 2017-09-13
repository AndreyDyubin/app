package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
)

var S3 = new(s3.S3)

func ConnectS3(awsKey, awsSecret string) error {
	token := ""
	creds := credentials.NewStaticCredentials(awsKey, awsSecret, token)
	_, err := creds.Get()
	if err != nil {
		return err
	}
	cfg := aws.NewConfig().WithRegion("us-west-1").WithCredentials(creds)
	s, err := session.NewSession()
	if err != nil {
		return err
	}
	S3 = s3.New(s, cfg)
	return nil
}
