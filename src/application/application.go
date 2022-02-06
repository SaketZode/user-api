package application

import (
	"fmt"
)

func StartApplication() {
	migrateDatabase()
	mapURLS()
	fmt.Println("Starting Application Server....")
	router.Run(":8080")
}
