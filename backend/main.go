package main

import (
	"apitester/api"
	"apitester/database"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("API Tester is running...")
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Authorization"},
		AllowCredentials: true,
	}))

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
