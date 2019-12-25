package main

import (
	"go-id-bot/app/server"
)

func main() {
	s := server.Server{
		Version: "1",
		Secret: "DEV",
	}
	s.Run("127.0.0.1", 8080)
}
