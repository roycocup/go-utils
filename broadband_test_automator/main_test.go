package main

import (
	"testing"

	"fmt"

	"net/http/httptest"
	"strings"

	"github.com/sirupsen/logrus"
)

func TestGetUrl(t *testing.T) {
	logrus.Info("Test started")
	var url string
	url = "http://google.com"

	w := httptest.NewRecorder()
	r := httptest.NewRequest(
		"GET",
		"/v1/id_value",
		strings.NewReader(""),
	)

	resp, err := getUrl(url)

	if err != nil {
		logrus.WithError(err).Fatal("Problems")
		t.Fail()
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		//t.Fail()
	}
}
