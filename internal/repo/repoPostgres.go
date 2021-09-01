package repo

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-food-api/internal/food"
)

const (
	table = "food_info"
)

var HaveNotElementErr = errors.New("Have not element with same id")

type repoPostgres struct {
	db sqlx.DB
}

func (r *repoPostgres) AddEntities(foods []food.Food) error {
	for _, elem := range foods {
		err := r.AddEntity(elem)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *repoPostgres) AddEntity(food food.Food) error {
	query, args, err := sq.Insert(table).
		Columns("user_id", "type", "name", "portion_size").
		Values(food.UserId, food.Type, food.Name, food.PortionSize).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}
	res, err := r.db.Exec(query, args...)

	if err != nil {
		return err
	}
	if ra, rerr := res.RowsAffected(); rerr != nil {
		return rerr
	} else if ra < 1 {
		return HaveNotElementErr
	}
	return nil
}

func (r *repoPostgres) ListEntities(limit, offset uint64) ([]food.Food, error) {
	query, args, err := sq.Select("*").
		From(table).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	sliceFoods := []food.Food{}
	for rows.Next() {
		tmpFood := food.Food{}
		err = rows.Scan(&tmpFood.Id, &tmpFood.UserId, &tmpFood.Type, &tmpFood.Name, &tmpFood.PortionSize)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, HaveNotElementErr
			}
			return nil, err
		}
		sliceFoods = append(sliceFoods, tmpFood)
	}
	return sliceFoods, nil
}
func (r *repoPostgres) DescribeEntity(foodId uint64) (*food.Food, error) {
	query, args, err := sq.Select("id", "user_id", "type", "name", "portion_size").
		From(table).
		Where(sq.Eq{"id": foodId}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	row := r.db.QueryRow(query, args...)
	f := food.Food{}
	err = row.Scan(&f.Id, &f.UserId, &f.Type, &f.Name, &f.PortionSize)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, HaveNotElementErr
		}
		return nil, err
	}
	return &f, nil
}
func (r *repoPostgres) RemoveEntity(foodId uint64) error {
	query, args, err := sq.Delete(table).
		Where(sq.Eq{"id": foodId}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	res, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	if ra, rerr := res.RowsAffected(); rerr != nil {
		return rerr
	} else if ra < 1 {
		return HaveNotElementErr
	}
	return nil
}
