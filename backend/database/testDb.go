package database

import (
	"apitester/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&models.Test{})
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDatabase() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func GetTestByID(id int) (*models.Test, error) {
	var test models.Test
	result := DB.First(&test, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &test, nil
}

// func GetTestByIDAsync(id int) <-chan AsyncTestResult {
// 	resultChan := make(chan AsyncTestResult)

// 	go func() {
// 		defer close(resultChan)

// 		test, err := GetTestByID(id)

// 		result := AsyncTestResult{
// 			Test:  test,
// 			Error: err,
// 		}

// 		resultChan <- result
// 	}()

// 	return resultChan
// }

func CreateTest(test *models.Test) error {
	result := DB.Create(test)
	return result.Error
}

func UpdateTest(test *models.Test) error {
	result := DB.Save(test)
	return result.Error
}

type AsyncTestResult struct {
	Test  *models.Test
	Error error
}
