package config

import (
	"github.com/iniyusril/presence-api/util"
	"github.com/joho/godotenv"
)

func InitializedEnvironment() {
	err := godotenv.Load(".env")
	util.PanicIfError(err)
}
