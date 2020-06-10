package connection

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// NewCache create cache(redis) connection
func NewCache(host, pass string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pass,
		DB:       0,
	})
}

// NewLog create loggin instance or connection to centralized logging server
func NewLog(service string) *logrus.Entry {
	log := logrus.New()

	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	l := log.WithFields(logrus.Fields{
		"service-name": service,
	})

	return l
}

// NewDB create database sql connection to remote server
func NewDB(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "dsn was wrong")
	}

	return db, nil
}
