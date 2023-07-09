package main

import "os"

type config struct {
	port           string
	structuredLogs bool
}

func (c config) Port() string {
	return c.port
}

// get all the configurations for this project
//
// NOTE: there are probably good libs in golang to handle this,
// but it's not necessary yet given the scope of the project
func newConfig() config {

	port := os.Getenv("PORT")
	structuredLogs := os.Getenv("STRUCTURED_LOGS")
	c := config{
		port:           "8080",
		structuredLogs: false,
	}

	if port != "" {
		c.port = port
	}

	//if "1" == structuredLogs {
	//	c.structuredLogs = true
	//}

	return c
}
