package database

import (
	"apitester/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type TestRepository struct {
	db *gorm.DB
}

func NewTestRepository(db *gorm.DB) *TestRepository {
	return &TestRepository{db: db}
}

func InitDatabase() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Test{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *TestRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *TestRepository) CloseDatabase() error {
	if r.db == nil {
		return nil
	}
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (r *TestRepository) GetTestByID(id int) (*models.Test, error) {
	var test models.Test
	result := r.db.First(&test, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &test, nil
}

func (r *TestRepository) GetAllTests() ([]models.Test, error) {
	var tests []models.Test
	result := r.db.Find(&tests)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return tests, nil // No tests found
	}

	return tests, nil
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

func (r *TestRepository) CreateTest(test *models.Test) error {
	result := r.db.Create(test)
	return result.Error
}

func (r *TestRepository) UpdateTest(test *models.Test) error {
	result := r.db.Save(test)
	return result.Error
}
