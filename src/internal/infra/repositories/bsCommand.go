package BsRepository

import (
	"context"
	"database/sql"
	"errors"
	BeerStyleEntity "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/domain/beerStyle"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

type DbCmd struct {
	Db *sql.DB
}

func NewCmd(db *sql.DB) *DbCmd {
	return &DbCmd{Db: db}
}

func (db *DbCmd) Create(
	ctx context.Context,
	beerStyle *BeerStyleEntity.BeerStyle,
) error {
	query := `
		INSERT INTO beer_styles (id, name, active, min_temp, max_temp, type_temp)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := db.Db.ExecContext(
		ctx,
		query,
		beerStyle.ID,
		beerStyle.Name,
		beerStyle.Active,
		beerStyle.MinTemp,
		beerStyle.MaxTemp,
		beerStyle.TempType,
	)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return ErrBeerStyleAlreadyExists
			}
		}
		return err
	}

	return nil
}

func (db *DbCmd) Update(
	ctx context.Context,
	id string,
	name *string,
	minTemp *float64,
	maxTemp *float64,
) error {

	setClauses := []string{}
	args := []any{}
	argPos := 1

	if name != nil {
		setClauses = append(setClauses, "name = $"+strconv.Itoa(argPos))
		args = append(args, *name)
		argPos++
	}

	if minTemp != nil {
		setClauses = append(setClauses, "min_temp = $"+strconv.Itoa(argPos))
		args = append(args, *minTemp)
		argPos++
	}

	if maxTemp != nil {
		setClauses = append(setClauses, "max_temp = $"+strconv.Itoa(argPos))
		args = append(args, *maxTemp)
		argPos++
	}

	query := `
		UPDATE beer_styles
		SET ` + strings.Join(setClauses, ", ") + `
		WHERE id = $` + strconv.Itoa(argPos)

	args = append(args, id)

	res, err := db.Db.ExecContext(ctx, query, args...)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return ErrBeerStyleAlreadyExists
			}
		}
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (db *DbCmd) SetActive(
	ctx context.Context,
	id string,
	active bool,
) error {
	query := `
		UPDATE beer_styles
		SET active = $1
		WHERE id = $2
	`

	res, err := db.Db.ExecContext(ctx, query, active, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
