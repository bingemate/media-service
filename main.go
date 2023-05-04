package main

import (
	"flag"
	"github.com/bingemate/media-service/cmd"
	"github.com/bingemate/media-service/initializers"
	"log"
)

// @title Media Service API
// @description This is the API for the Media Service application
// @description This help to give info about the media files and metadata
// @description This also help to manage the media files for admins
func main() {
	flag.Parse()
	env, err := initializers.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	logFile := initializers.InitLog(env.LogFile)
	defer logFile.Close()
	log.Println("Starting server mode...")
	cmd.Serve(env)
}
