package helper

import (
	"time"
	"xsis-academy-test-service-movie/constant"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type responseError struct {
	Code            int                  `json:"code"`
	Title           string               `json:"title"`
	UserMessage     constant.UserMessage `json:"user_message"`
	InternalMessage string               `json:"internal_message"`
	Time            string               `json:"time"`
}

type responseSuccess struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Time string      `json:"time"`
}

var location *time.Location

func init() {
	location, _ = time.LoadLocation(constant.TimeLocation)
}

func HttpSimpleResponse(c *fiber.Ctx, httpStatus int) error {
	return c.Status(httpStatus).SendString(fasthttp.StatusMessage(httpStatus))
}

func HttpResponseError(c *fiber.Ctx, code constant.InternalError, internalMessage string) error {
	return c.Status(code.Info().HttpCode).JSON(&responseError{
		Code:            int(code),
		Title:           code.Info().Title,
		UserMessage:     code.Info().UserMessage,
		InternalMessage: internalMessage,
		Time:            time.Now().In(location).Format(time.RFC3339),
	})
}

func HttpResponseSuccess(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(&responseSuccess{
		Code: fiber.StatusOK,
		Data: data,
		Time: time.Now().In(location).Format(time.RFC3339),
	})
}

func HttpResponseFileSuccess(c *fiber.Ctx, fullPathFile string) error {
	return c.Status(fiber.StatusOK).SendFile(fullPathFile, false)
}
