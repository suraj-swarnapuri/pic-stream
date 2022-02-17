package main

import "github.com/suraj-swarnapuri/pic-stream/backend"

func main() {
	router := backend.SetupServer()
	router.Run("localhost:8080")
}
