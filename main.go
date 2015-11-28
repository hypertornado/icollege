package main

import (
	"github.com/hypertornado/icollege/server"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

func main() {
	app := kingpin.New("icollege", "icollege")

	serverCommand := app.Command("server", "Run server")
	port := serverCommand.Flag("port", "server port").Default("8580").Short('p').Int()
	developmentMode := serverCommand.Flag("development", "Is in development mode").Default("false").Short('d').Bool()

	command, err := app.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	switch command {

	case serverCommand.FullCommand():
		server.Start(*port, *developmentMode)
	}
}
