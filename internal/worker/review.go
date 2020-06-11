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
	sqlGetReviewByID         = `SELECT comment FROM review WHERE review_id=?`
	sqlCheckKeywordIsExist   = `SELECT 1 FROM food_dictionary WHERE name=?`
	sqlSearchReviewByKeyword = `SELECT name FROM food_dictionary WHERE name LIKE '%?%'`

	cacheReviewKey  = `review:%v`
	cacheKeywordKey = `keyword:%v`
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

func (r *Review) GetReviewInCache(id string) (string, error) {
	key := fmt.Sprintf(cacheReviewKey, id)

	v, err := r.Cache.Get(context.TODO(), key).Result()
	if err != nil {
		return "", errors.Wrap(err, "redis get")
	}

	return v, nil
}

func (r *Review) GetReviewInDB(id string) (string, error) {
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

func (r *Review) SearchReviewByKeywordInDB(keyword string) ([]string, error) {
	stmt, err := r.DB.PrepareContext(context.TODO(), sqlSearchReviewByKeyword)
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	result, err := stmt.QueryContext(context.TODO(), keyword)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}
	defer result.Close()

	var reviews []string
	for result.Next() {
		var review string
		if err := result.Scan(&review); err != nil {
			return nil, errors.Wrap(err, "scan")
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (r *Review) SearchKeywordInCache(keyword string) (string, error) {
	key := fmt.Sprintf(cacheKeywordKey, keyword)

	v, err := r.Cache.Get(context.TODO(), key).Result()
	if err != nil {
		return "", errors.Wrap(err, "redis get")
	}

	return v, nil
}

func (r *Review) SearchKeywordInDB(keyword string) error {
	stmt, err := r.DB.PrepareContext(context.TODO(), sqlCheckKeywordIsExist)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	var result int
	if err := stmt.QueryRowContext(context.TODO(), keyword).Scan(&result); err != nil {
		return errors.Wrap(err, "query row")
	}
	if result == 0 {
		return errors.New("keyword not exists")
	}

	return nil
}
