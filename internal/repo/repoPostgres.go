package repo

import (
	"context"
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

func (r *repoPostgres) AddEntities(ctx context.Context, foods []food.Food) error {
	query := sq.Insert(table).
		Columns("user_id", "type", "name", "portion_size").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, elem := range foods {
		query = query.Values(elem.UserId, elem.Type, elem.Name, elem.PortionSize)
	}

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (r *repoPostgres) AddEntity(ctx context.Context, food food.Food) error {
	query, args, err := sq.Insert(table).
		Columns("user_id", "type", "name", "portion_size").
		Values(food.UserId, food.Type, food.Name, food.PortionSize).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
func (r *repoPostgres) ListEntities(ctx context.Context, limit, offset uint64) ([]food.Food, error) {
	query, args, err := sq.Select("*").
		From(table).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	sliceFoods := []food.Food{}
	defer rows.Close()
	for rows.Next() {
		tmpFood := food.Food{}
		err = rows.Scan(&tmpFood.Id, &tmpFood.UserId, &tmpFood.Type, &tmpFood.Name, &tmpFood.PortionSize)
		if err != nil {
			return nil, err
		}
		sliceFoods = append(sliceFoods, tmpFood)
	}
	return sliceFoods, rows.Err()
}
func (r *repoPostgres) DescribeEntity(ctx context.Context, foodId uint64) (*food.Food, error) {
	query, args, err := sq.Select("id", "user_id", "type", "name", "portion_size").
		From(table).
		Where(sq.Eq{"id": foodId}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	row := r.db.QueryRowContext(ctx, query, args...)
	f := food.Food{}
	err = row.Scan(&f.Id, &f.UserId, &f.Type, &f.Name, &f.PortionSize)

	if err != nil {
		return nil, err
	}
	return &f, nil
}
func (r *repoPostgres) RemoveEntity(ctx context.Context, foodId uint64) error {
	query, args, err := sq.Delete(table).
		Where(sq.Eq{"id": foodId}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repoPostgres) UpdateEntity(ctx context.Context, food food.Food) error {
	query, args, err := sq.Update(table).
		SetMap(map[string]interface{}{"id": food.Id, "user_id": food.UserId,
			"type": food.Type, "name": food.Name, "portion_size": food.PortionSize}).
		Where(sq.Eq{"id": food.Id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}

	return nil
}
func (r *repoPostgres) MultiAddEntity(ctx context.Context, foods [][]food.Food) error {
	for _, elem := range foods {
		err := r.AddEntities(ctx, elem)
		if err != nil {
			return err
		}
	}
	return nil
}
