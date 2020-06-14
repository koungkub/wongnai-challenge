package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber"
	"github.com/koungkub/wongnai-challenge/internal/model"
	"github.com/koungkub/wongnai-challenge/internal/service"
	"github.com/koungkub/wongnai-challenge/internal/worker"
)

var (
	cacheTimeout = time.Hour
)

// GetReviewByID handler get review by specific review_id
func GetReviewByID() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		id, conf := c.Params("id"), c.Locals("conf").(*model.Conf)
		w := worker.NewReview(conf.DB, conf.Cache)
		svc := service.NewPreview(w)

		review, err := svc.GetReviewInCache(context.TODO(), id)
		if err != nil {
			review, err = svc.GetReviewInDB(context.TODO(), id)
			if err != nil || len(review) == 0 {
				c.Status(fiber.StatusUnprocessableEntity).Render("review", fiber.Map{
					"Content": "review not found",
				})
				return
			}
			svc.SetReviewInCache(context.TODO(), id, review, cacheTimeout)
		}

		c.Status(fiber.StatusOK).Render("review", fiber.Map{
			"Content": review,
		})
	}
}

// SearchReviewByQuery handler get review by specific review_keyword
func SearchReviewByQuery() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		keyword, conf := c.Query("query"), c.Locals("conf").(*model.Conf)
		w := worker.NewReview(conf.DB, conf.Cache)
		svc := service.NewPreview(w)

		if _, err := svc.SearchKeywordInCache(context.TODO(), keyword); err != nil {
			if err := svc.SearchKeywordInDB(context.TODO(), keyword); err != nil {
				c.Status(fiber.StatusUnprocessableEntity).Render("keyword", fiber.Map{
					"Content": "keyword not found",
				})
				return
			}
			svc.SetKeywordInCache(context.TODO(), keyword, "1", cacheTimeout)
		}

		review, err := svc.SearchReviewByKeywordInDB(context.TODO(), keyword)
		if err != nil {
			c.Status(fiber.StatusUnprocessableEntity).Render("keyword", fiber.Map{
				"Content": "review not found",
			})
			return
		}

		c.Status(fiber.StatusOK).Render("keyword", fiber.Map{
			"Reviews": review,
			"Keyword": keyword,
		})
	}
}

// EditReview handler edit review by specific review_id
func EditReview() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		id, review, conf := c.Params("id"), new(model.Review), c.Locals("conf").(*model.Conf)
		w := worker.NewReview(conf.DB, conf.Cache)
		svc := service.NewPreview(w)

		if err := c.BodyParser(review); err != nil {
			c.Status(fiber.StatusUnprocessableEntity).Render("edit", fiber.Map{
				"Content": "can not parse request body",
			})
		}

		rows, err := svc.EditReviewInDB(context.TODO(), id, review.Data.Comment)
		if err != nil || rows <= 0 {
			c.Status(fiber.StatusUnprocessableEntity).Render("edit", fiber.Map{
				"Content": "review not change because review_id not found or comment have no difference",
			})
			return
		}

		if err := svc.DelReviewKey(context.TODO(), id); err != nil {
			c.Status(fiber.StatusInternalServerError).Render("edit", fiber.Map{
				"Content": "server error",
			})
			return
		}

		c.Status(fiber.StatusOK).Render("edit", fiber.Map{
			"Content": "review updated",
		})
	}
}
