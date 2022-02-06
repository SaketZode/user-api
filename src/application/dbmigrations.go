package application

import (
	"fmt"
	"user-api/databaseconnection/postgresdb"
	"user-api/models/usermodels"
)

func migrateDatabase() {
	fmt.Println("Running migrations....")
	err := postgresdb.DatabaseConnection.AutoMigrate(&usermodels.User{})
	if err != nil {
		fmt.Println("Error in migrating database:", err)
		panic("Database Migration Failed!!!!")
	}
}
