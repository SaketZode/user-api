package postgresdb

import (
	"fmt"
	"os"
	"user-api/databaseconnection"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DatabaseConnection *gorm.DB
)

func init() {
	var pgconnection *postgresdbconnection = readEnvVariables()
	DatabaseConnection = pgconnection.ConnectDb()
}

func readEnvVariables() (pgConn *postgresdbconnection) {
	dbHost := os.Getenv(databaseconnection.DB_HOST)
	dbPort := os.Getenv(databaseconnection.DB_PORT)
	schemaName := os.Getenv(databaseconnection.DB_SCHEMA)
	username := os.Getenv(databaseconnection.DB_USERNAME)
	password := os.Getenv(databaseconnection.DB_PASSWORD)

	pgConn = &postgresdbconnection{
		DbHost:     dbHost,
		DbPort:     dbPort,
		SchemaName: schemaName,
		Username:   username,
		Password:   password,
	}

	return pgConn
}

type postgresdbconnection struct {
	DbHost     string
	DbPort     string
	SchemaName string
	Username   string
	Password   string
}

func (pgConn *postgresdbconnection) GetConnectionString() (connString string) {
	connString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", pgConn.DbHost, pgConn.Username, pgConn.Password, pgConn.SchemaName, pgConn.DbPort)
	return connString
}

func (pgConn *postgresdbconnection) ConnectDb() *gorm.DB {
	connectionString := pgConn.GetConnectionString()
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		fmt.Println("Error at opening DB connection:", err)
		panic("Failed to establish postgres DB connection!!")
	}

	fmt.Println("DB connection established!!")
	return db
}
