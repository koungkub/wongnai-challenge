package worker

import (
	"context"
	"time"
)

// Reviewer interface for review worker
type Reviewer interface {
	GetReviewInCache(ctx context.Context, id string) (string, error)
	GetReviewInDB(ctx context.Context, id string) (string, error)
	SetReviewInCache(ctx context.Context, id string, review string, exp time.Duration) error

	SearchKeywordInCache(ctx context.Context, keyword string) (string, error)
	SearchKeywordInDB(ctx context.Context, keyword string) error
	SetKeywordInCache(ctx context.Context, keyword string, review string, exp time.Duration) error

	SearchReviewByKeywordInDB(ctx context.Context, keyword string) ([]string, error)

	EditReviewInDB(ctx context.Context, id, review string) (int64, error)
	DelReviewKey(ctx context.Context, id string) error
}
