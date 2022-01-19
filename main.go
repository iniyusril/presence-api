package main

import "github.com/iniyusril/presence-api/config"

func main() {
	config.InitializedEnvironment()
	router := InitializedServer()
	router.Logger.Fatal(router.Start(":9001"))
}
