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
	sqlSearchReviewByKeyword = `SELECT comment FROM review WHERE comment LIKE ?`
	sqlEditReviewByID        = `UPDATE review SET comment=? WHERE review_id=?`

	cacheReviewKey  = `review:%v`
	cacheKeywordKey = `keyword:%v`
)

// Review worker for review logic
type Review struct {
	DB    *sql.DB
	Cache *redis.Client
}

// NewReview get review instance
func NewReview(db *sql.DB, cache *redis.Client) Reviewer {
	return &Review{
		DB:    db,
		Cache: cache,
	}
}

// GetReviewInCache get review by id in cache
func (r *Review) GetReviewInCache(id string) (string, error) {
	key := fmt.Sprintf(cacheReviewKey, id)

	v, err := r.Cache.Get(context.TODO(), key).Result()
	if err != nil {
		return "", errors.Wrap(err, "redis get")
	}

	return v, nil
}

// GetReviewInDB get review by id in database
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

// SetReviewInCache set review by id in cache
func (r *Review) SetReviewInCache(id string, review string, exp time.Duration) error {
	key := fmt.Sprintf(cacheReviewKey, id)

	err := r.Cache.Set(context.TODO(), key, review, exp).Err()
	if err != nil {
		return errors.Wrap(err, "redis set")
	}

	return nil
}

// SearchReviewByKeywordInDB search review by keyword in database
func (r *Review) SearchReviewByKeywordInDB(keyword string) ([]string, error) {
	stmt, err := r.DB.PrepareContext(context.TODO(), sqlSearchReviewByKeyword)
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	queryArgs := fmt.Sprintf("%%%v%%", keyword)
	rows, err := stmt.QueryContext(context.TODO(), queryArgs)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}
	defer rows.Close()

	var reviews []string
	for rows.Next() {
		var review string
		if err := rows.Scan(&review); err != nil {
			return nil, errors.Wrap(err, "scan")
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

// SearchKeywordInCache search keyword in cache
func (r *Review) SearchKeywordInCache(keyword string) (string, error) {
	key := fmt.Sprintf(cacheKeywordKey, keyword)

	v, err := r.Cache.Get(context.TODO(), key).Result()
	if err != nil {
		return "", errors.Wrap(err, "redis get")
	}

	return v, nil
}

// SearchKeywordInDB search keyword in database
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

// SetKeywordInCache set keyword in cache
func (r *Review) SetKeywordInCache(keyword string, review string, exp time.Duration) error {
	key := fmt.Sprintf(cacheKeywordKey, keyword)

	err := r.Cache.Set(context.TODO(), key, review, exp).Err()
	if err != nil {
		return errors.Wrap(err, "redis set")
	}

	return nil
}

// EditReviewInDB edit review by id in database
func (r *Review) EditReviewInDB(id, review string) (int64, error) {
	stmt, err := r.DB.PrepareContext(context.TODO(), sqlEditReviewByID)
	if err != nil {
		return 0, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(context.TODO(), review, id)
	if err != nil {
		return 0, errors.Wrap(err, "exec")
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rows affected")
	}

	return rows, nil
}

// DelReviewKey delete key in cache
func (r *Review) DelReviewKey(id string) error {
	key := fmt.Sprintf(cacheReviewKey, id)

	err := r.Cache.Del(context.TODO(), key).Err()
	if err != nil {
		return errors.Wrap(err, "redis del")
	}

	return nil
}
