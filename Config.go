package main

//Config comment
type Config struct {
	BaseURL       string
	RedisEndpoint string
	RedisDatabase string
	RedisPass     string
	Database      string
}

var c Config
