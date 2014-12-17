package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostToS3(t *testing.T) {
	ts := serveHTTP(t)
	s3URL = ts.URL
	postToS3("foobar.txt", []byte("file data in here"))
}

func TestGenerateSignature(t *testing.T) {
	Convey("A proper signature should be generated", t, func() {
		date := "Mon, 1 Jan 1970 15:04:05 +0000"
		tropoAWSKey = "abc123"
		tropoAWSSecret = "def456"
		tropoAWSBucket = "foobar"
		So(generateSignature("foobar.txt", date), ShouldEqual, "J2Y442YmYgu/x4YEOYjEpzRfs8s=")
	})
}

func TestComputeHmac1(t *testing.T) {
	Convey("A proper signature should be generated", t, func() {
		So(computeHmac1("foobar", "secret"), ShouldEqual, "f1wOnLLwcTexwCSRCNXEAKPDm+U=")
	})
}

// serveHTTP serves up a test server emulating AWS
func serveHTTP(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		req.Body.Close()

		Convey("The body should be correct", t, func() {
			So(string(body), ShouldEqual, "file data in here")
		})

		Convey("All of the headers should be set", t, func() {
			// So(req.Header["Host"][0], ShouldEqual, s3Domain)
			So(req.Header["Date"][0], ShouldNotEqual, "")
			So(req.Header["Content-Type"][0], ShouldEqual, "application/x-compressed-tar")
			So(req.Header["Authorization"][0], ShouldNotEqual, "")
		})
	}))
}
