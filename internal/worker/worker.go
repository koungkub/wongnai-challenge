package worker

import "time"

type Reviewer interface {
	GetReviewInCache(id string) (string, error)
	GetReviewInDB(id string) (string, error)
	SetReviewInCache(id string, review string, exp time.Duration) error
	SearchReviewByKeywordInDB(keyword string) ([]string, error)
	SearchKeywordInCache(keyword string) (string, error)
	SearchKeywordInDB(keyword string) error
}
