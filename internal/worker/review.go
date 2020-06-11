package worker

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

const (
	sqlGetReviewByID = `SELECT comment FROM review WHERE review_id=?`

	cacheReviewKey = `review:%v`
)

type Review struct {
	DB    *sql.DB
	Cache *redis.Client
}

func NewReview(db *sql.DB, cache *redis.Client) Reviewer {
	return &Review{
		DB:    db,
		Cache: cache,
	}
}

func (r *Review) GetReviewByCache(id string) (string, error) {
	key := fmt.Sprintf(cacheReviewKey, id)

	v, err := r.Cache.Get(context.TODO(), key).Result()
	if err != nil {
		return "", errors.Wrap(err, "redis get")
	}

	return v, nil
}

func (r *Review) GetReviewByDB(id string) (string, error) {
	stmt, err := r.DB.PrepareContext(context.TODO(), sqlGetReviewByID)
	if err != nil {
		return "", errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	var review string
	if err := stmt.QueryRowContext(context.TODO(), id).Scan(&review); err != nil {
		return "", errors.Wrap(err, "query row")
	}

	return review, nil
}

func (r *Review) SetReviewInCache(id string, review string, exp time.Duration) error {
	key := fmt.Sprintf(cacheReviewKey, id)

	err := r.Cache.Set(context.TODO(), key, review, exp).Err()
	if err != nil {
		return errors.Wrap(err, "redis set")
	}

	return nil
}
