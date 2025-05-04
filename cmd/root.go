package cmd

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"gophinder/internal/handler"
	"gophinder/internal/notifier"
	"gophinder/utils"
)

var templateFile string
var showHelp bool
var portNumber int
var configFilePath string

func Execute() {
	flag.StringVar(&templateFile, "t", "templates/index.html", "Specify a path to a template file")
	flag.StringVar(&configFilePath, "config", "config.json", "Path to JSON config file for notifications")
	flag.IntVar(&portNumber, "p", 8080, "Web Server port")
	flag.BoolVar(&showHelp, "h", false, "Show usage")
	flag.Parse()

	if showHelp {
		flag.Usage()
		return
	}

	config, err := utils.LoadConfig(configFilePath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	notificationService, err := notifier.NewNotifierService(*config)
	if err != nil {
		log.Fatalf("Error initializing notification service: %v", err)
	}

	http.HandleFunc("/", handler.ServeTemplate(templateFile, notificationService))
	http.HandleFunc("/track", handler.HandleTrack(notificationService))
	http.HandleFunc("/login", handler.HandleLogin(notificationService))

	server := fmt.Sprintf("http://localhost:%d", portNumber)
	fmt.Printf("\033[32m[+] Server running on %s\033[0m\n", server)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", portNumber), nil))
}
