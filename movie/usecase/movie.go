package usecase

import (
	"context"
	"xsis-academy-test-service-movie/domain"
	"xsis-academy-test-service-movie/helper"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type movieUseCase struct {
	movieUseCase   domain.MovieUseCase
	movieMySQLRepo domain.MovieMySQLRepo
}

func NewMovieUsecase(MovieMySQLRepo domain.MovieMySQLRepo) domain.MovieUseCase {
	return &movieUseCase{
		movieMySQLRepo: MovieMySQLRepo,
	}
}

func (mvu *movieUseCase) PostMovie(ctx context.Context, request domain.RequestMovie) (err error) {
	parentPath := viper.GetString("server.url_assets")
	subPath := "images/banner"
	imagePath, err := helper.SaveImageToLocalDrive(request.Image, parentPath, subPath)
	if err != nil {
		return err
	}

	request.ImagePath = imagePath
	err = mvu.movieMySQLRepo.PostMovie(ctx, request)
	if err != nil {
		log.Error(err)
		return err
	}
	return
}

func (mvu *movieUseCase) GetAllMovie(ctx context.Context, request domain.RequestParamMovie) (response domain.ResponseGetAllMovie, err error) {
	resMovieCount, err := mvu.movieMySQLRepo.CountDataMovie(ctx, request)
	if err != nil {
		return domain.ResponseGetAllMovie{}, err
	}

	resMovie, err := mvu.movieMySQLRepo.GetAllMovie(ctx, request)
	if err != nil {
		return response, err
	}

	response = domain.ResponseGetAllMovie{
		MetaData: resMovieCount,
		Data:     resMovie,
	}

	log.Info(resMovie)
	return
}

func (mvu *movieUseCase) DeleteMovie(ctx context.Context, id int) (err error) {
	_, err = mvu.movieMySQLRepo.GetDetailMovie(ctx, id)
	if err != nil {
		return err
	}

	err = mvu.movieMySQLRepo.DeleteMovie(ctx, id)
	if err != nil {
		return err
	}
	return
}

func (mvu *movieUseCase) GetDetailMovie(ctx context.Context, id int) (response domain.ResponseMovie, err error) {
	response, err = mvu.movieMySQLRepo.GetDetailMovie(ctx, id)
	if err != nil {
		return response, err
	}
	return
}

func (mvu *movieUseCase) UpdateMovie(ctx context.Context, id int, request domain.RequestMovie) (err error) {
	_, err = mvu.movieMySQLRepo.GetDetailMovie(ctx, id)
	if err != nil {
		return err
	}

	if request.Image.Filename != "" {
		parentPath := viper.GetString("server.url_assets")
		subPath := "images/banner"
		imagePath, err := helper.SaveImageToLocalDrive(request.Image, parentPath, subPath)
		if err != nil {
			return err
		}
		request.ImagePath = imagePath
	}
	err = mvu.movieMySQLRepo.UpdateMovie(ctx, id, request)
	if err != nil {
		log.Error(err)
		return err
	}
	return
}
