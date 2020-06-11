package worker

import "time"

type Reviewer interface {
	GetReviewInCache(id string) (string, error)
	GetReviewInDB(id string) (string, error)
	SetReviewInCache(id string, review string, exp time.Duration) error
}
