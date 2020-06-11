package handler

import (
	"time"

	"github.com/gofiber/fiber"
	"github.com/koungkub/wongnai/internal/model"
	"github.com/koungkub/wongnai/internal/worker"
)

// GetReviewByID handler get review by specific review_id
func GetReviewByID() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		id := c.Params("id")
		conf := c.Locals("conf").(*model.Conf)
		r := worker.NewReview(conf.DB, conf.Cache)

		review, err := r.GetReviewByCache(id)
		if err != nil {
			review, err = r.GetReviewByDB(id)
			if err != nil || len(review) == 0 {
				c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"message": "id not match any reviews",
				})
				return
			}
			r.SetReviewInCache(id, review, time.Hour)
		}

		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": review,
		})
	}
}

// SearchReviewByQuery handler get review by specific review_keyword
func SearchReviewByQuery() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {

	}
}

// EditReview handler edit review by specific review_id
func EditReview() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {

	}
}
