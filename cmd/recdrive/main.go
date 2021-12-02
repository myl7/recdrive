package main

import (
	"bufio"
	"flag"
	"github.com/myl7/recdrive"
	"io/ioutil"
	"log"
	"net/http"
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
			names = append(names, recdrive.Filename(items[i]))
		}
		println(strings.Join(names, " "))
	case action == "cp" || action == "copy":
		dest := flag.Arg(2)
		if dest == "" {
			log.Fatalln("destination is required")
		}

		if strings.HasPrefix(dest, ":") {
			// Upload
			dest = strings.TrimPrefix(dest, ":")

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

			id, err := drive.Upload(dest, file, filename, filesize)
			if err != nil {
				log.Fatalln(err.Error())
			}

			println("uploaded ok to", id)
		} else if strings.HasPrefix(path, ":") {
			// Download
			path = strings.TrimPrefix(path, ":")

			s, err := drive.Download(path)
			if err != nil {
				log.Fatalln(err.Error())
			}

			file, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Fatalln(err.Error())
			}

			res, err := http.Get(s)
			if err != nil {
				log.Fatalln(err.Error())
			}

			if res.StatusCode != 200 {
				b, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Fatalln(err.Error())
				}

				log.Fatalf("fetch file from %s failed: status %d reason %s\n", s, res.StatusCode, string(b))
			}

			_, err = bufio.NewReader(res.Body).WriteTo(file)
			if err != nil {
				log.Fatalln(err.Error())
			}

			err = file.Close()
			if err != nil {
				log.Fatalln(err.Error())
			}

			println("download ok")
		} else {
			log.Fatalln("both path and dest are not a remote file")
		}
	default:
		log.Fatalln("unknown action", action)
	}
}
