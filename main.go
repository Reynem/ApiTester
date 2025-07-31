package main

import (
	"apitester/api"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("API Tester is running...")
	e := echo.New()
	testManagerGroup := e.Group("/api")
	api.TestManager(testManagerGroup)

	if err := e.Start(":8080"); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	} else {
		fmt.Println("Server started successfully on port 8080")
	}
}
