package mysql

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"xsis-academy-test-service-movie/domain"

	"github.com/labstack/gommon/log"
)

type mysqlMovieRepository struct {
	Conn *sql.DB
}

func NewMySQLMovieRepository(Conn *sql.DB) domain.MovieMySQLRepo {
	return &mysqlMovieRepository{Conn}
}

func (db *mysqlMovieRepository) PostMovie(ctx context.Context, request domain.RequestMovie) (err error) {
	query := `INSERT INTO movie (title, description, rating, image, dtm_crt, dtm_upd)
              VALUES (?, ?, ?, ?, NOW(), NOW())`

	_, err = db.Conn.ExecContext(ctx, query, request.Title, request.Description, request.Rating, request.ImagePath)

	if err != nil {
		return err
	}

	return
}

func (db *mysqlMovieRepository) CountDataMovie(ctx context.Context, request domain.RequestParamMovie) (response domain.MetaData, err error) {
	var query string
	query = "SELECT COUNT(id) as total FROM movie WHERE 1=1"
	var limit, page int
	var order string
	var args []interface{}

	if request.Search != nil {
		query += " AND title LIKE ? OR description LIKE ? OR rating LIKE ?"
		args = append(args, "%"+*request.Search+"%", "%"+*request.Search+"%", "%"+*request.Search+"%")
	}

	if request.Limit != nil {
		limit = *request.Limit
	}

	if request.Page != nil {
		page = *request.Page
	}

	if request.Order != nil {
		order = *request.Order
		query += " ORDER BY " + order
	}

	if limit > 0 {
		query += " LIMIT ? OFFSET ?"
		args = append(args, limit, (page-1)*limit)
	}

	log.Debug(query)

	var count int
	err = db.Conn.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return response, err
	}

	if count == 0 {
		err = errors.New("Not found")
		return domain.MetaData{}, err
	}

	totalPage := uint(count) / uint(*request.Limit)
	if uint(count)%uint(*request.Limit) != 0 {
		totalPage++
	}

	response = domain.MetaData{
		TotalData: uint(count),
		TotalPage: totalPage,
		Page:      uint(page),
		Limit:     uint(*request.Limit),
	}

	return response, nil
}

func (db *mysqlMovieRepository) UpdateMovie(ctx context.Context, id int, request domain.RequestMovie) (err error) {
	query := `UPDATE movie
              SET title = ?, description = ?, rating = ?, image = ?, dtm_upd = NOW()
              WHERE id = ?`

	_, err = db.Conn.ExecContext(ctx, query, request.Title, request.Description, request.Rating, request.ImagePath, id)

	if err != nil {
		return err
	}

	return nil
}

func (db *mysqlMovieRepository) GetAllMovie(ctx context.Context, request domain.RequestParamMovie) (response []domain.ResponseMovie, err error) {
	query := `SELECT id, title, description, rating, image, dtm_crt, dtm_upd FROM movie`
	var limit, page int
	var order string
	var args []interface{}

	if request.Search != nil {
		query += " WHERE title LIKE ? OR description LIKE ? OR rating LIKE ?"
		args = append(args, "%"+*request.Search+"%", "%"+*request.Search+"%", "%"+*request.Search+"%")
	}

	if request.Order != nil {
		order = *request.Order
		query += " ORDER BY " + order
	}
	if request.Page != nil {
		page = *request.Page
	}

	if request.Limit != nil {
		limit = *request.Limit
		if limit > 0 {
			query += " LIMIT ? OFFSET ?"
			args = append(args, limit, (page-1)*limit)
		}
	}

	log.Debug(query)
	rows, err := db.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	var movies []domain.ResponseMovie
	for rows.Next() {
		var i domain.ResponseMovie
		var dtmCrt, dtmUpd time.Time
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Rating,
			&i.Image,
			&dtmCrt,
			&dtmUpd,
		); err != nil {
			log.Error(err)
			return nil, err
		}

		i.DtmCrt = dtmCrt.Format("2006-01-02 15:04:05")
		i.DtmUpd = dtmUpd.Format("2006-01-02 15:04:05")

		movies = append(movies, i)
	}

	return movies, nil
}

func (db *mysqlMovieRepository) DeleteMovie(ctx context.Context, id int) (err error) {
	query := `DELETE FROM movie WHERE id = ?`
	_, err = db.Conn.ExecContext(ctx, query, id)
	if err != nil {
		log.Error(err)
		return err
	}

	return
}

func (db *mysqlMovieRepository) GetDetailMovie(ctx context.Context, id int) (response domain.ResponseMovie, err error) {
	query := `SELECT id, title, description, rating, image, dtm_crt, dtm_upd FROM movie WHERE id = ?`

	row := db.Conn.QueryRowContext(ctx, query, id)
	var dtmCrt, dtmUpd time.Time
	err = row.Scan(
		&response.ID,
		&response.Title,
		&response.Description,
		&response.Rating,
		&response.Image,
		&dtmCrt,
		&dtmUpd,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New("Not found")
			return domain.ResponseMovie{}, err
		}
		log.Error(err)
		return domain.ResponseMovie{}, err
	}

	response.DtmCrt = dtmCrt.Format("2006-01-02 15:04:05")
	response.DtmUpd = dtmUpd.Format("2006-01-02 15:04:05")

	return response, nil
}
