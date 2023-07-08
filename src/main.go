package main

import "github.com/pintjuk/routemaster/src/http_resources"

func main() {
	config := newConfig()
	http_resources.StartHttpServer(config, nil)
}
