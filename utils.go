package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration
var logger *log.Logger

// Convenience function for printing to stdout
func p(a ...interface{}) {
	fmt.Println(a)
}

func init() {
	loadConfig()
	file, err := os.OpenFile("flare.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

// for logging
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// version
func version() string {
	return "0.1"
}
