package http

import (
	"xsis-academy-test-service-movie/movie/delivery/http/handler"
	// "xsis-academy-test-service-movie/delivery/http/handler"
	"xsis-academy-test-service-movie/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// RouterAPI is the main router for this Service Insurance REST API
func RouterAPI(app *fiber.App, MovieUseCase domain.MovieUseCase) {
	handlerMovie := &handler.MovieHandler{MovieUseCase: MovieUseCase}
	basePath := viper.GetString("server.base_path")

	movie := app.Group(basePath)

	movie.Use(cors.New(cors.Config{
		AllowOrigins: viper.GetString("middleware.allows_origin"),
	}))

	log.Info(handlerMovie)
	// Public API Route
	movie.Get("/movie", handlerMovie.GetAllMovie)
	movie.Post("/movie", handlerMovie.PostMovie)
	movie.Delete("/movie/:id", handlerMovie.DeleteMovie)
	movie.Patch("/movie/:id", handlerMovie.UpdateMovie)
	movie.Get("/movie/:id", handlerMovie.GetDetailMovie)

}
