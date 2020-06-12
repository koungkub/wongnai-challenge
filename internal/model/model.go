package model

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

// Conf store connection
type Conf struct {
	DB    *sql.DB
	Log   *logrus.Entry
	Cache *redis.Client
}

// NewConf create configuration instance
func NewConf(db *sql.DB, log *logrus.Entry, cache *redis.Client) *Conf {
	return &Conf{
		DB:    db,
		Log:   log,
		Cache: cache,
	}
}
