package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	AWS_S3_REGION = "us-east-2"
	AWS_S3_BUCKET = "polymer-long-tail-seo"
)

func main() {
	uploadFilename := os.Args[1]

	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
	if err != nil {
		fmt.Println("error in session creation!", err)
	}

	err = uploadFile(session, uploadFilename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("[*] uplading file: %v -> bucket: %v // SUCCESS!\n", uploadFilename, AWS_S3_BUCKET)
	}
}

func uploadFile(session *session.Session, uploadFilename string) error {
	file, err := os.Open(uploadFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var fileSize int64 = fileInfo.Size()

	fileBuffer := make([]byte, fileSize)
	file.Read(fileBuffer)

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(AWS_S3_BUCKET),
		Key:                  aws.String(uploadFilename),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}
