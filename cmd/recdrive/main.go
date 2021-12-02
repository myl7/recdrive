package main

import (
	"flag"
	"github.com/myl7/recdrive"
	"log"
	"os"
	"path/filepath"
	"strings"
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

		var names []string
		for i := range items {
			names = append(names, items[i].Name)
		}
		println(strings.Join(names, " "))
	case action == "cp" || action == "copy":
		filename := filepath.Base(path)
		if filename == "" {
			log.Fatalln("invalid path")
		}

		file, err := os.Open(path)
		if err != nil {
			log.Fatalln(err.Error())
		}

		stat, err := file.Stat()
		if err != nil {
			log.Fatalln(err.Error())
		}

		filesize := stat.Size()

		dest := flag.Arg(2)
		if dest == "" {
			log.Fatalln("destination is required")
		}

		id, err := drive.Upload(dest, file, filename, filesize)
		if err != nil {
			log.Fatalln(err.Error())
		}

		println("uploaded ok to", id)
	default:
		log.Fatalln("unknown action", action)
	}
}
