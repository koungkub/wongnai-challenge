package handler

import (
	"time"

	"github.com/gofiber/fiber"
	"github.com/koungkub/wongnai/internal/model"
	"github.com/koungkub/wongnai/internal/worker"
)

var (
	cacheTimeout = time.Hour
)

// GetReviewByID handler get review by specific review_id
func GetReviewByID() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		id := c.Params("id")
		conf := c.Locals("conf").(*model.Conf)
		r := worker.NewReview(conf.DB, conf.Cache)

		review, err := r.GetReviewInCache(id)
		if err != nil {
			review, err = r.GetReviewInDB(id)
			if err != nil || len(review) == 0 {
				c.Status(fiber.StatusUnprocessableEntity).Render("review", fiber.Map{
					"Content": "review not found",
				})
				return
			}
			r.SetReviewInCache(id, review, cacheTimeout)
		}

		c.Status(fiber.StatusOK).Render("review", fiber.Map{
			"Content": review,
		})
	}
}

// SearchReviewByQuery handler get review by specific review_keyword
func SearchReviewByQuery() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		keyword := c.Query("query")
		conf := c.Locals("conf").(*model.Conf)
		r := worker.NewReview(conf.DB, conf.Cache)

		if _, err := r.SearchKeywordInCache(keyword); err != nil {
			if err := r.SearchKeywordInDB(keyword); err != nil {
				c.Status(fiber.StatusUnprocessableEntity).Render("keyword", fiber.Map{
					"Content": "keyword not found",
				})
				return
			}
			r.SetKeywordInCache(keyword, "1", cacheTimeout)
		}

		review, err := r.SearchReviewByKeywordInDB(keyword)
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
		id, review := c.Params("id"), new(model.Review)
		conf := c.Locals("conf").(*model.Conf)
		r := worker.NewReview(conf.DB, conf.Cache)

		if err := c.BodyParser(review); err != nil {
			c.Status(fiber.StatusUnprocessableEntity).Render("edit", fiber.Map{
				"Content": "can not parse request body",
			})
		}

		rows, err := r.EditReviewInDB(id, review.Data.Comment)
		if err != nil || rows <= 0 {
			c.Status(fiber.StatusUnprocessableEntity).Render("edit", fiber.Map{
				"Content": "review not change because review_id not found or comment have no difference",
			})
			return
		}
		// purge review_id

		c.Status(fiber.StatusOK).Render("edit", fiber.Map{
			"Content": "review updated",
		})
	}
}
