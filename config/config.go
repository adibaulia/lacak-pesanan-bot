package config

import "os"

var (
	PORT      = os.Getenv("PORT")       // sets automatically
	PublicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
	TELETOKEN = os.Getenv("TOKEN")      // you must add it to your config vars
)

func init() {

}

func GetInstance() {}
