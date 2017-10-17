package main

import (
	"net/url"
)

type Config struct {
	BaseUrl  string
	Database string
}

var c Config

func (c Config) baseUrl() string {

	u, _ := url.Parse(c.BaseUrl)

	return u.Host
}
