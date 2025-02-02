package routes

import (
	"time"
	"github.com/gofiber/fiber/v2"
)

type request struct {
	URL			string			`json:"url"`
	CustomShort string			`json:"short"`
	Expiry		time.Duration	`json:"expiry"`
}

type response sturct {
	URL				string			`json:"url"`
	CustomShort 	string			`json:"customshort"`
	Expiry			time.Duration	`json:"expiry"`
	XRateRemaining	int				`json:"rate_limit"`
	XRateLimitRest	time.Duration	`json:"rate_limit_reset"`
}


func ShortenURL(c *fiber.Ctx) error {

	body := new(request)

	if err := c.BodyParse(&body); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse JSON"})
	}


	// implement rate limiting
	//
	//
	// //chjeck if the input if an actual URL

	if !govalidator.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid URL"})
	}

	// check for domain error

	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid URL"})
	}
	// enforce https
}
