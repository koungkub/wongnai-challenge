package service

import (
	"context"
	"time"

	"github.com/koungkub/wongnai/internal/worker"
)

type Preview struct {
	Worker worker.Reviewer
}

func NewPreview(w worker.Reviewer) Servicer {
	return &Preview{
		Worker: w,
	}
}

// GetReviewInCache get review by id in cache
func (p *Preview) GetReviewInCache(ctx context.Context, id string) (string, error) {
	v, err := p.Worker.GetReviewInCache(ctx, id)
	if err != nil {
		return "", err
	}

	return v, nil
}

// GetReviewInDB get review by id in database
func (p *Preview) GetReviewInDB(ctx context.Context, id string) (string, error) {
	v, err := p.Worker.GetReviewInDB(ctx, id)
	if err != nil {
		return "", err
	}

	return v, nil
}

// SetReviewInCache set review by id in cache
func (p *Preview) SetReviewInCache(ctx context.Context, id string, review string, exp time.Duration) error {
	err := p.Worker.SetReviewInCache(ctx, id, review, exp)
	if err != nil {
		return err
	}

	return nil
}

// SearchReviewByKeywordInDB search review by keyword in database
func (p *Preview) SearchReviewByKeywordInDB(ctx context.Context, keyword string) ([]string, error) {
	v, err := p.Worker.SearchReviewByKeywordInDB(ctx, keyword)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// SearchKeywordInCache search keyword in cache
func (p *Preview) SearchKeywordInCache(ctx context.Context, keyword string) (string, error) {
	v, err := p.Worker.SearchKeywordInCache(ctx, keyword)
	if err != nil {
		return "", err
	}

	return v, nil
}

// SearchKeywordInDB search keyword in database
func (p *Preview) SearchKeywordInDB(ctx context.Context, keyword string) error {
	err := p.Worker.SearchKeywordInDB(ctx, keyword)
	if err != nil {
		return err
	}

	return nil
}

// SetKeywordInCache set keyword in cache
func (p *Preview) SetKeywordInCache(ctx context.Context, keyword string, review string, exp time.Duration) error {
	err := p.Worker.SetKeywordInCache(ctx, keyword, review, exp)
	if err != nil {
		return err
	}

	return nil
}

// EditReviewInDB edit review by id in database
func (p *Preview) EditReviewInDB(ctx context.Context, id, review string) (int64, error) {
	v, err := p.Worker.EditReviewInDB(ctx, id, review)
	if err != nil {
		return 0, err
	}

	return v, nil
}

// DelReviewKey delete key in cache
func (p *Preview) DelReviewKey(ctx context.Context, id string) error {
	err := p.Worker.DelReviewKey(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
