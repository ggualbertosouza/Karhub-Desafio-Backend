package BsRepository

import (
	"context"
	"database/sql"
	"errors"
)

type DbQuery struct {
	Db *sql.DB
}

func NewQuery(db *sql.DB) *DbQuery {
	return &DbQuery{Db: db}
}

func (q *DbQuery) GetById(ctx context.Context, id string) (*BsModel, error) {
	query := `
	SELECT id, name, min_temp, max_temp, type_temp, active
	FROM beer_styles
	WHERE id = $1
	`

	var model BsModel
	err := q.Db.
		QueryRowContext(ctx, query, id).
		Scan(&model.Id, &model.Name, &model.MinTemp, &model.MaxTemp, &model.TempType, &model.Active)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrBeerStyleNotFound
		}
		return nil, err
	}

	return &model, nil
}

func (q *DbQuery) GetByName(ctx context.Context, name string) (*BsModel, error) {
	query := `
		SELECT id, name, min_temp, max_temp, type_temp, active
		FROM beer_styles
		WHERE name = $1
	`

	var model BsModel
	err := q.Db.
		QueryRowContext(ctx, query, name).
		Scan(&model.Id, &model.Name, &model.MinTemp, &model.MaxTemp, &model.TempType, &model.Active)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrBeerStyleNotFound
		}
		return nil, err
	}

	return &model, nil
}

func (q *DbQuery) ListAll(ctx context.Context) ([]BsModel, error) {
	query := `
		SELECT id, name, min_temp, max_temp, type_temp, active
		FROM beer_styles
		WHERE active = true
		ORDER BY name ASC
	`

	rows, err := q.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []BsModel

	for rows.Next() {
		var model BsModel
		if err := rows.Scan(
			&model.Id,
			&model.Name,
			&model.MinTemp,
			&model.MaxTemp,
			&model.TempType,
			&model.Active,
		); err != nil {
			return nil, err
		}

		list = append(list, model)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
