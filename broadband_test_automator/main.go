package main

import (
	"net/http"

	"../debug"
	"github.com/sirupsen/logrus"
)

const siteUrl = "http://www.speedtest.net/"

func main() {
	getUrl(siteUrl)
}

func getUrl(url string) (resp *http.Response, err error) {
	resp, err = http.Get(url)
	if err != nil {
		logrus.WithError(err).Fatal("Unable to get site")
	}
	return
}

func sfsd() {
	debug.DD("sdfsdf")
}
