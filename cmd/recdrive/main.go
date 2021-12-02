package main

import (
	"flag"
	"github.com/myl7/recdrive"
	"log"
)

func main() {
	token := flag.String("token", "", "authentication token for recdrive")
	flag.Parse()

	if *token == "" {
		log.Fatalln("token is required")
	}

	action := flag.Arg(0)
	if action == "" {
		log.Fatalln("action is required")
	}

	path := flag.Arg(1)
	if path == "" {
		log.Fatalln("path is required")
	}

	drive := recdrive.NewDrive(recdrive.Options{AuthToken: *token})

	switch {
	case action == "ls" || action == "list":
		items, err := drive.List(path)
		if err != nil {
			log.Fatalln(err.Error())
		}

		for i := range items {
			print(items[i].Name, " ")
		}
	default:
		log.Fatalln("unknown action", action)
	}
}
