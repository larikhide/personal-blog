package main

import (
	"flag"
	"fmt"
)

func main() {
	flagConfigFile := flag.String("config", "./config.yaml", "Flag config in yaml format")
	flag.Parse()

	config, err := ReadConfig(*flagConfigFile)
	if err != nil {
		panic(fmt.Sprintf("Cannot read config file %s", err))
	}

	lg, err := ConfigureLogger(&config.Logger)
	if err != nil {
		panic(fmt.Sprintf("Cannot configure logger %s", err))
	}

	server := NewServer(&config.Server, lg)
	if err := server.Start(); err != nil {
		lg.WithError(err).Fatal(("Fatal error. Cannot start server"))
	}
}
