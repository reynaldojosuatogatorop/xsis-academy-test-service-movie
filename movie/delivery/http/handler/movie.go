package handler

import (
	"strconv"
	"xsis-academy-test-service-movie/domain"
	"xsis-academy-test-service-movie/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

type MovieHandler struct {
	MovieUseCase   domain.MovieUseCase
	MovieMySQLRepo domain.MovieMySQLRepo
}

func (mh *MovieHandler) GetAllMovie(c *fiber.Ctx) error {
	var input domain.RequestParamMovie
	search := c.Query("search")
	if search != "" {
		input.Search = &search
	}

	limit := c.Query("limit")
	if limit == "" {
		limitInt := viper.GetInt("database.default_limit_query")
		log.Info("Parameter limit not set, running with default limit query from config", limitInt)

		input.Limit = &limitInt
	} else {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
		}
		input.Limit = &limitInt
	}

	page := c.Query("page")
	if page == "" {
		pageInt := viper.GetInt("database.default_page")
		log.Info("Parameter limit not set, running with default limit query from config", pageInt)

		input.Page = &pageInt
	} else {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
		}
		input.Page = &pageInt
	}

	order := c.Query("order")
	if order == "" {
		input.Order = nil
	} else {
		input.Order = &order
	}
	res, err := mh.MovieUseCase.GetAllMovie(c.Context(), input)
	if err != nil {
		return c.SendStatus(fasthttp.StatusInternalServerError)
	}

	return c.Status(fasthttp.StatusOK).JSON(res)
}

func (mh *MovieHandler) PostMovie(c *fiber.Ctx) (err error) {
	var input domain.RequestMovie
	err = c.BodyParser(&input)
	if err != nil {
		log.Error(err.Error())
		return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	gambarBinary, err := c.FormFile("image")
	if err != nil {
		log.Error(err)
		return helper.HttpSimpleResponse(c, fasthttp.StatusInternalServerError)
	}

	if gambarBinary == nil {
		return c.Status(fasthttp.StatusBadRequest).SendString("No upload images")
	}

	input.Image = *gambarBinary

	ratingFloat, err := strconv.ParseFloat(input.Rating, 64)
	if err != nil {
		log.Error("Error converting string to float64:", err)
		return
	}

	input.FloatRating = ratingFloat
	err = mh.MovieUseCase.PostMovie(c.Context(), input)
	if err != nil {
		log.Error(err)
		return helper.HttpSimpleResponse(c, fasthttp.StatusInternalServerError)
	}
	return c.SendStatus(fasthttp.StatusCreated)
}

func (mh *MovieHandler) UpdateMovie(c *fiber.Ctx) (err error) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Error(err)
		return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
	}
	var input domain.RequestMovie
	err = c.BodyParser(&input)
	if err != nil {
		log.Error(err.Error())
		return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	gambarBinary, err := c.FormFile("image")
	if err != nil {
		log.Error(err)
		return helper.HttpSimpleResponse(c, fasthttp.StatusInternalServerError)
	}

	// if gambarBinary == nil {
	// 	return c.Status(fasthttp.StatusBadRequest).SendString("No upload images")
	// }

	if gambarBinary != nil {
		input.Image = *gambarBinary

	}

	ratingFloat, err := strconv.ParseFloat(input.Rating, 64)
	if err != nil {
		log.Error("Error converting string to float64:", err)
		return
	}

	input.FloatRating = ratingFloat
	err = mh.MovieUseCase.UpdateMovie(c.Context(), int(id), input)
	if err != nil {
		log.Error(err)
		return helper.HttpSimpleResponse(c, fasthttp.StatusInternalServerError)
	}
	return c.Status(fasthttp.StatusOK).SendString("Updated")
}

func (mh *MovieHandler) DeleteMovie(c *fiber.Ctx) (err error) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Error(err)
		return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	err = mh.MovieUseCase.DeleteMovie(c.Context(), int(id))
	if err != nil {
		if err.Error() == "Not found" {
			return helper.HttpSimpleResponse(c, fasthttp.StatusNotFound)
		}
		return err
	}
	return c.Status(fasthttp.StatusOK).SendString("Deleted")
}

func (mh *MovieHandler) GetDetailMovie(c *fiber.Ctx) (err error) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Error(err)
		return helper.HttpSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	res, err := mh.MovieUseCase.GetDetailMovie(c.Context(), int(id))
	if err != nil {
		if err.Error() == "Not found" {
			return helper.HttpSimpleResponse(c, fasthttp.StatusNotFound)
		}
		return err
	}
	return c.Status(fasthttp.StatusOK).JSON(res)
}
