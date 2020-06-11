package worker

import "time"

type Reviewer interface {
	GetReviewByCache(id string) (string, error)
	GetReviewByDB(id string) (string, error)
	SetReviewInCache(id string, review string, exp time.Duration) error
}
