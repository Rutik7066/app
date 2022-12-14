package credit

import (
	"github.com/gofiber/fiber/v2"
)

func GetPlan(c *fiber.Ctx) error {
	type Plan struct {
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Credit   int     `json:"credit"`
		Validity string  `json:"validity"`
	}
	type AutoGenerated struct {
		Plan []Plan `json:"plan"`
	}
	plan := AutoGenerated{
		Plan: []Plan{
			{
				Name:     "250",
				Price:    250.00,
				Credit:   250,
				Validity: "30 Days",
			}, {
				Name:     "500",
				Price:    500.00,
				Credit:   500,
				Validity: "30 Days",
			}, {
				Name:     "1000",
				Price:    1000.00,
				Credit:   1500,
				Validity: "30 Days",
			}, {
				Name:     "3000",
				Price:    3000.00,
				Credit:   5000,
				Validity: "30 Days",
			}, {
				Name:     "5000",
				Price:    5000.00,
				Credit:   10000,
				Validity: "90 Days",
			},
		},
	}

	return c.Status(fiber.StatusOK).JSON(&plan)
}
