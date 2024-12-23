// Package cs/v1 provides all functionality to interract with YourBestie
// in-house applicatin: Cloud Storage API.
package v1

import (
	"crypto/tls"
	"net/http"
)

type Service = *service

type service struct {
	Options
	http *http.Client
}

type Options struct {
	BaseURL     string
	AccessToken string
}

func Setup(opt Options) Service {
	return &service{
		Options: opt,
		http: &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}},
	}
}
