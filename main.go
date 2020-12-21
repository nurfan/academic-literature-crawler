package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nurfan/academic-literature-crawler/route"

	"github.com/joho/godotenv"
	logger "github.com/sirupsen/logrus"
)

func init() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		logger.Fatalf("Error loading .env file")
	}

	Formatter := new(logger.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	logger.SetFormatter(Formatter)
}

func main() {
	e := route.Init()
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Panic(fmt.Sprint(err))
	}

	fmt.Println(string(data))
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
