package db

import (
	"github.com/Yamiyo/account/glob/config"
)

// Database ...
type Database interface {
	Connect(config config.Databases) error
	Size() int
	Close()
}
