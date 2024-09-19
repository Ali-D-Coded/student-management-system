package app

import (
	"log"
	"os"
	"student-management-system/database"
	"student-management-system/router"
	"student-management-system/utils"

	"student-management-system/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Indexable struct {
	collection string
	fields     []string
}


func SetupAndRunApp() error {


	Indexables := []Indexable{
	   {
		   collection: "users",
		   fields:     []string{"username", "email", "refreshToken"},
	   },
	}
	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// start database
	err = database.StartMongoDB()
	if err != nil {
		return err
	}

	for _, indexable := range Indexables {
		// Create unique indexes for the users collection
		err = utils.CreateUniqueIndexes(indexable.collection, indexable.fields...)
		if err != nil {
			log.Fatal("Failed to create unique indexes for users:", err)
		}
	}



	// defer closing database
	defer database.CloseMongoDB()

	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	

	
	// setup routes
	router.SetupRoutes(app)



	// get the port and start
	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}
