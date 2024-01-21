package domain

import (
	"context"
	"mime/multipart"
)

type RequestMovie struct {
	Title       string               `json:"title" form:"title"`
	Description string               `json:"description" form:"description"`
	Rating      string               `json:"rating" form:"rating"`
	Image       multipart.FileHeader `json:"gambar" form:"gambar"`
	ImagePath   string               `json:"image_path"`
	FloatRating float64              `json:"float_rating"`
}

type ResponseMovie struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Image       string  `json:"image"`
	DtmCrt      string  `json:"dtm_crt"`
	DtmUpd      string  `json:"dtm_upd"`
}

type RequestParamMovie struct {
	Page   *int    `json:"page"`
	Limit  *int    `json:"limit"`
	Order  *string `json:"order"`
	Search *string `json:"search"`
}

type ResponseGetAllMovie struct {
	MetaData MetaData        `json:"meta_data"`
	Data     []ResponseMovie `json:"data"`
}

type MetaData struct {
	TotalData uint   `json:"total_data"`
	TotalPage uint   `json:"total_page"`
	Page      uint   `json:"page"`
	Limit     uint   `json:"limit"`
	Sort      string `json:"sort"`
	Order     string `json:"order"`
}

type MovieUseCase interface {
	PostMovie(ctx context.Context, request RequestMovie) error
	GetAllMovie(ctx context.Context, request RequestParamMovie) (response ResponseGetAllMovie, err error)
	DeleteMovie(ctx context.Context, id int) (err error)
	UpdateMovie(ctx context.Context, id int, request RequestMovie) (err error)
	GetDetailMovie(ctx context.Context, id int) (response ResponseMovie, err error)
}

type MovieMySQLRepo interface {
	PostMovie(ctx context.Context, request RequestMovie) error
	CountDataMovie(ctx context.Context, request RequestParamMovie) (response MetaData, err error)
	GetAllMovie(ctx context.Context, request RequestParamMovie) (response []ResponseMovie, err error)
	DeleteMovie(ctx context.Context, id int) (err error)
	UpdateMovie(ctx context.Context, id int, request RequestMovie) (err error)
	GetDetailMovie(ctx context.Context, id int) (response ResponseMovie, err error)
}

type MovieGRPCRepo interface {
}
