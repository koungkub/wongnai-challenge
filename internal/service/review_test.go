package service

import (
	"context"
	"errors"
	"testing"
	"time"

	_mocks "github.com/koungkub/wongnai-challenge/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDelReviewKey(t *testing.T) {
	t.Run("del_review_key_success", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("DelReviewKey", context.TODO(), "1").Return(nil).Once()

		svc := NewPreview(w)
		err := svc.DelReviewKey(context.TODO(), "1")

		assert.NoError(t, err)
		w.AssertExpectations(t)
	})

	t.Run("del_review_key_failure", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("DelReviewKey", context.TODO(), "1").Return(errors.New("err")).Once()

		svc := NewPreview(w)
		err := svc.DelReviewKey(context.TODO(), "1")

		assert.Error(t, err)
		w.AssertExpectations(t)
	})
}

func TestEditReviewInDB(t *testing.T) {
	t.Run("edit_review_in_db_success", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("EditReviewInDB", context.TODO(), "pizza", "pizza is not delicious").Return(int64(1), nil).Once()

		svc := NewPreview(w)
		result, err := svc.EditReviewInDB(context.TODO(), "pizza", "pizza is not delicious")

		assert.NoError(t, err)
		assert.Equal(t, int64(1), result)
		w.AssertExpectations(t)
	})

	t.Run("edit_review_in_db_failure", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("EditReviewInDB", context.TODO(), "pizza", "").Return(int64(1), errors.New("err")).Once()

		svc := NewPreview(w)
		_, err := svc.EditReviewInDB(context.TODO(), "pizza", "")

		assert.Error(t, err)
		w.AssertExpectations(t)
	})
}

func TestSearchReviewByKeywordInDB(t *testing.T) {
	t.Run("search_review_by_keyword_in_db_success", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("SearchReviewByKeywordInDB", context.TODO(), "pizza").Return([]string{"good", "great"}, nil).Once()

		svc := NewPreview(w)
		reviews, err := svc.SearchReviewByKeywordInDB(context.TODO(), "pizza")

		assert.NoError(t, err)
		assert.Equal(t, []string{"good", "great"}, reviews)
		w.AssertExpectations(t)
	})

	t.Run("search_review_by_keyword_in_db_failure", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("SearchReviewByKeywordInDB", context.TODO(), "").Return([]string{}, errors.New("err")).Once()

		svc := NewPreview(w)
		_, err := svc.SearchReviewByKeywordInDB(context.TODO(), "")

		assert.Error(t, err)
		w.AssertExpectations(t)
	})
}

func TestSetKeywordInCache(t *testing.T) {
	t.Run("set_keyword_in_cache_success", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("SetKeywordInCache", context.TODO(), "pizza", "pizza is good", time.Hour).Return(nil).Once()

		svc := NewPreview(w)
		err := svc.SetKeywordInCache(context.TODO(), "pizza", "pizza is good", time.Hour)

		assert.NoError(t, err)
		w.AssertExpectations(t)
	})

	t.Run("set_keyword_in_cache_failure", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("SetKeywordInCache", context.TODO(), "pizza", "", time.Hour).Return(errors.New("err")).Once()

		svc := NewPreview(w)
		err := svc.SetKeywordInCache(context.TODO(), "pizza", "", time.Hour)

		assert.Error(t, err)
		w.AssertExpectations(t)
	})
}

func TestSearchKeywordInDB(t *testing.T) {
	t.Run("search_keyword_in_db_success", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("SearchKeywordInDB", context.TODO(), "pizza").Return(nil).Once()

		svc := NewPreview(w)
		err := svc.SearchKeywordInDB(context.TODO(), "pizza")

		assert.NoError(t, err)
		w.AssertExpectations(t)
	})

	t.Run("search_keyword_in_db_failure", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("SearchKeywordInDB", context.TODO(), "pizza").Return(errors.New("err")).Once()

		svc := NewPreview(w)
		err := svc.SearchKeywordInDB(context.TODO(), "pizza")

		assert.Error(t, err)
		w.AssertExpectations(t)
	})
}

func TestSearchKeywordInCache(t *testing.T) {
	t.Run("search_keyword_in_cache_success", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("SearchKeywordInCache", context.TODO(), "pizza").Return("1", nil).Once()

		svc := NewPreview(w)
		keyword, err := svc.SearchKeywordInCache(context.TODO(), "pizza")

		assert.NoError(t, err)
		assert.Equal(t, "1", keyword)

		w.AssertExpectations(t)
	})

	t.Run("search_keyword_in_cache_failure", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("SearchKeywordInCache", context.TODO(), "").Return("", errors.New("err")).Once()

		svc := NewPreview(w)
		_, err := svc.SearchKeywordInCache(context.TODO(), "")

		assert.Error(t, err)
		w.AssertExpectations(t)
	})
}

func TestSetReviewInCache(t *testing.T) {
	t.Run("set_review_in_cache_success", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("SetReviewInCache", context.TODO(), "1", "that good", time.Hour).Return(nil).Once()

		svc := NewPreview(w)
		err := svc.SetReviewInCache(context.TODO(), "1", "that good", time.Hour)

		assert.NoError(t, err)
		w.AssertExpectations(t)
	})

	t.Run("set_review_in_cache_failure", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("SetReviewInCache", context.TODO(), "1", "", time.Hour).Return(errors.New("err")).Once()

		svc := NewPreview(w)
		err := svc.SetReviewInCache(context.TODO(), "1", "", time.Hour)

		assert.Error(t, err)
		w.AssertExpectations(t)
	})
}

func TestGetReviewInDB(t *testing.T) {
	t.Run("get_review_in_db_success", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("GetReviewInDB", context.TODO(), "1").Return("delicious", nil).Once()

		svc := NewPreview(w)
		result, err := svc.GetReviewInDB(context.TODO(), "1")

		assert.NoError(t, err)
		assert.Equal(t, "delicious", result)
		w.AssertExpectations(t)
	})

	t.Run("get_review_in_db_failure", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("GetReviewInDB", context.TODO(), "1").Return("", errors.New("err")).Once()

		svc := NewPreview(w)
		_, err := svc.GetReviewInDB(context.TODO(), "1")

		assert.Error(t, err)
		w.AssertExpectations(t)
	})
}

func TestGetReviewInCache(t *testing.T) {
	t.Run("get_review_in_cache_success", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("GetReviewInCache", context.TODO(), "1").Return("delicious", nil).Once()

		svc := NewPreview(w)
		result, err := svc.GetReviewInCache(context.TODO(), "1")

		assert.NoError(t, err)
		assert.Equal(t, "delicious", result)
		w.AssertExpectations(t)
	})

	t.Run("get_review_in_cache_failure", func(t *testing.T) {
		w := new(_mocks.Reviewer)
		w.On("GetReviewInCache", context.TODO(), "1").Return("", errors.New("err")).Once()

		svc := NewPreview(w)
		_, err := svc.GetReviewInCache(context.TODO(), "1")

		assert.Error(t, err)
		w.AssertExpectations(t)
	})
}
