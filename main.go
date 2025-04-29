package main

import (
	"log"
	"os"

	"services-arraya-attendance/db"
	"services-arraya-attendance/forms"
	"services-arraya-attendance/routes"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	//Load the .env file
	if os.Getenv("ENV") != "PRODUCTION" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("error: failed to load the env file")
		}
	} else if os.Getenv("ENV_TYPE") == "CUSTOM" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("error: failed to load the env file")
		}
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Custom form validator
	binding.Validator = new(forms.DefaultValidator)

	//Start PostgreSQL database
	//Example: db.GetDB() - More info in the models folder
	db.Init()

	//Start Redis on database 1 - it's used to store the JWT but you can use it for anythig else
	//Example: db.GetRedis().Set(KEY, VALUE, at.Sub(now)).Err()
	db.InitRedis(1)

	//Start the default gin server
	routes.Run()

}
