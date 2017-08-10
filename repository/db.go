package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattes/migrate/driver/mysql"
	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
	"github.com/tokoku_backend/config"
)

type Manager interface {
	GetListOfWalmartCategory() (listOfWalmartCategory []WalmartCategory, err error)
	GetListOfWalmartItemByCategory(categoryID string, page int, perPage int) (listOfWalmartItem []WalmartItem, err error)

	AddWalmartItem(walmartItem *WalmartItem) (err error)

	DeleteWalmartItem(categoryID string) (err error)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func InitDBConnection() (err error) {
	url, err := config.Config.DB.ConnectionString()
	if err != nil {
		return
	}
	fmt.Println("opening db")
	db, err := gorm.Open("postgres", url)
	if err != nil {
		return
	}
	db.SingularTable(true)
	//db.LogMode(true)
	Mgr = &manager{db: db}

	return
}

//Migrate is needed to prepare a database tables
func Migrate(migrationPath string, num int) (errors []error) {
	url, err := config.Config.DB.MigrationString()
	if err != nil {
		return append(errors, err)
	}
	if num == 0 {
		errors, _ = migrate.UpSync(url, migrationPath)
	} else {
		errors, _ = migrate.MigrateSync(url, migrationPath, num)
	}
	return
}

//Drop is used for Clean up the database
func Drop(migrationPath string) (errors []error) {
	url, err := config.Config.DB.MigrationString()
	if err != nil {
		return append(errors, err)
	}
	errors, _ = migrate.DownSync(url, migrationPath)
	return
}
