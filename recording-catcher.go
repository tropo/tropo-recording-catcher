package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/pat"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	tropoAWSKey    string
	tropoAWSSecret string
	tropoAWSBucket string
	tropoPort      string
	s3Domain       string
	s3URL          string
)

func init() {
	// gets all of the environment variables to be set
	tropoAWSKey = os.Getenv("TROPO_AWS_KEY")
	if tropoAWSKey == "" {
		log.Fatal("Environment varialbe TROPO_AWS_KEY must be set.")
	}
	tropoAWSSecret = os.Getenv("TROPO_AWS_SECRET")
	if tropoAWSKey == "" {
		log.Fatal("Environment varialbe TROPO_AWS_SECRET must be set.")
	}
	tropoAWSBucket = os.Getenv("TROPO_AWS_BUCKET")
	if tropoAWSBucket == "" {
		log.Fatal("Environment varialbe TROPO_AWS_BUCKET must be set.")
	}
	tropoPort = os.Getenv("TROPO_PORT")
	if tropoPort == "" {
		log.Fatal("Environment varialbe TROPO_PORT must be set.")
	}
	s3Domain = tropoAWSBucket + ".s3.amazonaws.com"
	s3URL = "https://" + s3Domain
}

func main() {
	mux := pat.New()
	mux.Post("/recordings/{filename}", http.HandlerFunc(recordingHandler))
	mux.Put("/recordings/{filename}", http.HandlerFunc(recordingHandler))
	http.Handle("/", mux)
	log.Println("Listening on " + tropoPort + "...")
	http.ListenAndServe(":"+tropoPort, nil)
}

// handled the requests to post recordings to S3
func recordingHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	params := r.URL.Query()
	filename := params.Get(":filename")
	go postToS3(filename, body)
	return
}

// posts a the body sent to Amazon S3
func postToS3(filename string, body []byte) {
	url := s3URL + "/" + filename
	log.WithFields(log.Fields{
		"url": url,
	}).Info("Posting recording")

	client := &http.Client{}
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	setHeaders(req, filename)
	resp, err := client.Do(req)
	processResponse(resp, filename, err)
}

// setHeaders sets the required headers for S3
func setHeaders(req *http.Request, filename string) {
	date := time.Now().UTC().Format("Mon, _2 Jan 2006 15:04:05 +0000")
	signature := generateSignature(filename, date)
	req.Header.Add("Host", s3Domain)
	req.Header.Add("Date", date)
	req.Header.Add("Content-Type", "application/x-compressed-tar")
	req.Header.Add("Authorization", "AWS "+tropoAWSKey+":"+signature)
}

// generates the required signature for AWS
func generateSignature(filename string, date string) string {
	stringToSign := "PUT\n\napplication/x-compressed-tar\n" + date + "\n" + "/" + tropoAWSBucket + "/" + filename
	signature := computeHmac1(stringToSign, tropoAWSSecret)
	return signature
}

// computes the sha1 hmac and then base64 encodes
func computeHmac1(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// processes the reponses from AWS
func processResponse(resp *http.Response, filename string, err error) {
	if err != nil {
		logError(filename, err.Error())
	} else {
		if resp.StatusCode == 200 {
			log.WithFields(log.Fields{
				"filename": filename,
			}).Info("Successfully uploaded the recording")
		} else {
			logError(filename, resp.Status)
		}
	}
}

// logging function for response errors
func logError(filename string, err string) {
	log.WithFields(log.Fields{
		"filename": filename,
		"error":    err,
	}).Info("Failed to upload a recording")
}
