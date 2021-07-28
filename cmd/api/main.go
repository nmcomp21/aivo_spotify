package main

import "github.com/nmcomp/aivo-spotify/cmd/api/server"

func main() {
	server.New().Run(":8080")
}
