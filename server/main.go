package main

import (
	"flag"
	"log"

	weos "github.com/wepala/weos/controllers/rest"
	"golang.org/x/net/context"
)

var port = flag.String("port", "8681", "-port=8681")

func main() {
	flag.Parse()
	ctxt := context.Background()
	//instantiate weos with a reference to the OpenAPI specification
	api, err := weos.New("api.yaml")
	if err != nil {
		log.Fatalln("error loading api config", err)
	}
	err = api.Initialize(ctxt)
	if err != nil {
		log.Fatalln("error initializing api", err)
	}
	//start API
	api.EchoInstance().Logger.Fatal(api.EchoInstance().Start(":" + *port))
}
