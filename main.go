package main

import "hacktiv8-workshop-go-concurrency/router"

var (
	PORT = ":9090"
)

func main() {
	router.StartServer().Run(PORT)
}
