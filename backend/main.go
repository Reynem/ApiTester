package main

import (
	"apitester/api"
	"apitester/database"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("API Tester is running...")
	e := echo.New()
	testManagerGroup := e.Group("/api")

	sqliteDB, err := database.InitDatabase()
	if err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		return
	}

	db := database.NewTestRepository(sqliteDB)

	defer db.CloseDatabase()

	api.TestManager(testManagerGroup, db)

	if err := e.Start(":8080"); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	} else {
		fmt.Println("Server started successfully on port 8080")
	}
}
