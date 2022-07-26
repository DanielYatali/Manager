package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"manager/src"
	"os"
)

var port = flag.String("port", "8681", "-port=8681")

func main() {
	flag.Parse()
	ctxt := context.Background()
	//set default environment variables if none is set
	if os.Getenv("API_URL") == "" {
		err := os.Setenv("API_URL", "http://localhost:"+*port)
		if err != nil {
			fmt.Printf("error setting up '%s'", "API_URL")
		}
	}

	if os.Getenv("DB_DRIVER") == "" {
		err := os.Setenv("DB_DRIVER", "sqlite3")
		if err != nil {
			fmt.Printf("error setting up '%s'", "DB_DRIVER")
		}
	}

	if os.Getenv("DB_NAME") == "" {
		err := os.Setenv("DB_NAME", "bwcms.db")
		if err != nil {
			fmt.Printf("error setting up '%s'", "DB_NAME")
		}
	}

	api, _ := src.Init(ctxt, "./api.yaml")
	//start API
	api.EchoInstance().Logger.Fatal(api.EchoInstance().Start(":" + *port))
}
