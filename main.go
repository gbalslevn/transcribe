// https://www.jajaldoang.com/post/upload-file-to-aws-s3-with-go/

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// the uploader which can upload to S3
var uploader *s3manager.Uploader

func main() {
	uploader = NewUploader()
	uploadCode := time.Now().Unix()
	// upload - should take a file as parameter which will be uploaded to S3
	uploadFileToS3(uploadCode)
	//time.Sleep(10 * time.Second)
	// displays the transcription
	//displayTranscription(uploadCode)
}

// creates an uploader to the S3, to call the Upload method - https://github.com/aws/aws-sdk-go/
func NewUploader() *s3manager.Uploader {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	s3Config := &aws.Config{
		Region:      aws.String("eu-north-1"),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}
	s3Session := session.New(s3Config)
	uploader := s3manager.NewUploader(s3Session)
	return uploader
}

// uploads the file
func uploadFileToS3(uploadCode int64) {
	filename := "TranscribeTestLyd.m4a"
	// Find the last occurrence of "."
	lastDotIndex := strings.LastIndex(filename, ".")
	// Insert code before the last "."
	newFilename := strings.Join([]string{filename[:lastDotIndex], fmt.Sprintf("%d", uploadCode), filename[lastDotIndex:]}, "")
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	upInput := &s3manager.UploadInput{
		Bucket:      aws.String("audiofilebucketprojekt"),      // bucket's name
		Key:         aws.String("./audiofiles/" + newFilename), // files destination location - FILNAVN SKAL VÆRE LOWERCASE
		Body:        bytes.NewReader(file),                     // content of the file
		ContentType: aws.String("audio/wav"),                   // content type
	}
	res, err := uploader.UploadWithContext(context.Background(), upInput)
	if err == nil {
		log.Printf("Picture uploaded succesfully! res: %+v\n", res)
		
	} else {
		log.Printf("There was an error uploading: %+v\n", err)
	}
}

func displayTranscription(uploadCode int64) {
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	s3Config := &aws.Config{
		Region:      aws.String("eu-north-1"),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}
	s3Session := session.New(s3Config)
	// Create an S3 client using the session
	s3Client := s3.New(s3Session)

	// Create a GetObjectInput to retrieve the file from S3
	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String("audiofilebucketprojekt"),
		Key:    aws.String("transcribe-job-" + fmt.Sprint(uploadCode) + ".json"),
	}

	// Retrieve the transcription file from S3
	getObjectOutput, err := s3Client.GetObject(getObjectInput)
	if err != nil {
		log.Fatal(err)
	}
	defer getObjectOutput.Body.Close()

	// Read the contents of the file
	transcriptionBytes, err := ioutil.ReadAll(getObjectOutput.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Define a struct to represent the JSON data and to be able to read it
	type TranscriptsResponse struct {
		Results struct {
			Transcripts []struct {
				Transcript string `json:"transcript"`
			} `json:"transcripts"`
		} `json:"results"`
	}

	// Unmarshal the JSON data into the struct
	var result TranscriptsResponse
	err = json.Unmarshal(transcriptionBytes, &result)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	// Display the transcription to the user
	fmt.Println(result.Results.Transcripts[0].Transcript)
}
