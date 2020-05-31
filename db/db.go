package db

import (
	"errors"
	"fmt"
	"github.com/Yamiyo/account/glob/config"
	"strings"

	"github.com/Yamiyo/account/utils/log"
)

// Databases ...
type Databases map[string]Database

// Database Instances
var (
	MysqlInstance     = make(Databases)
	Close             func()
)

// Init ...
func Init(dbconf config.DatabaseConfig) error {
	for _, conf := range dbconf.Databases {
		tSlice := strings.Split(conf.Name, "_")

		if _, err := connect(conf, tSlice[0], tSlice[1], true); err != nil {
			log.Error(err)
			continue
		}
	}

	if len(MysqlInstance) <= 0 {
		return errors.New("no mysql database instance support")
	}

	Close = func() {
		for _, db := range MysqlInstance {
			db.Close()
		}
	}

	return nil
}

// Connect ...
func Connect(dbconf config.DatabaseConfig, brand string, dbType string) (Database, func(), error) {
	for _, conf := range dbconf.Databases {
		tSlice := strings.Split(conf.Name, "_")
		if tSlice[0] == brand && tSlice[1] == dbType {
			instance, err := connect(conf, tSlice[0], tSlice[1], false)
			if err != nil {
				return nil, nil, err
			}

			c := func() {
				instance.Close()
			}

			return instance, c, nil
		}
	}

	return nil, nil, fmt.Errorf("Connect fail, database configuration not found: %s_%s", brand, dbType)
}

func connect(conf config.Databases, brand string, dbType string, load bool) (Database, error) {
	var (
		err      error
		instance Database
	)

	switch dbType {
	case "mysql":
		instance, err = NewMySQL(conf)
		if err != nil {
			return nil, err
		}

		if load {
			MysqlInstance[brand] = instance
		}
	default:
		return nil, fmt.Errorf("Connect fail, unsupport database instance: %v", brand)
	}

	return instance, nil
}
