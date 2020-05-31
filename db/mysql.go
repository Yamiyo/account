package db

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	// Comment ...
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Yamiyo/account/glob/config"
	"github.com/Yamiyo/account/utils/log"
)

// MySQL ...
type MySQL struct {
	gorm *gorm.DB
}

// NewMySQL ...
func NewMySQL(config config.Databases) (Database, error) {
	db := new(MySQL)
	err := db.Connect(config)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetMySQL ...
func GetMySQL(databases Databases, brand string) (*MySQL, error) {
	database := databases[brand].(*MySQL)
	if database == nil {
		return nil, errors.New(strings.Join([]string{"error nil", brand, "mysql instance"}, " "))
	}
	return database, nil
}

// Connect ...
func (db *MySQL) Connect(config config.Databases) error {
	var err error
	connect := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		config.Username,
		config.Password,
		config.Address[0],
		config.DataBase,
	)

	db.gorm, err = gorm.Open("mysql", connect)
	if err != nil {
		return err
	}

	log.Infof("Database [%s] Connect success", config.Name)

	// TODO, load from config
	db.gorm.LogMode(false)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.gorm.DB().SetMaxIdleConns(50)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.gorm.DB().SetMaxOpenConns(5000)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.gorm.DB().SetConnMaxLifetime(15 * time.Minute)

	return nil
}

// Session ...
func (db *MySQL) Session() *gorm.DB {
	return db.gorm
}

// Begin ...
func (db *MySQL) Begin() *gorm.DB {
	return db.gorm.Begin()
}

// AutoMigrate ...
func (db *MySQL) AutoMigrate(model interface{}) error {
	return db.gorm.AutoMigrate(model).Error
}

// Close ...
func (db *MySQL) Close() {
	db.gorm.Close()
}

// Size ...
func (db *MySQL) Size() int {
	return 0
}
