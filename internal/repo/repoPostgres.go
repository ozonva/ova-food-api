package repo

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-food-api/internal/food"
)
const (
	table = "food_info"
)
type repoPostgres struct {
	db sqlx.DB
}

func (r *repoPostgres) AddEntities(foods []food.Food) error {
	for _,elem := range foods {
		query,args,err := sq.Insert(table).Columns("user_id","type","name","portion_size").
			Values(elem.UserId,elem.Type, elem.Name, elem.PortionSize).ToSql()
		if err != nil {
			return err
		}
		r.db.Exec(query,args)
	}

	return nil
}
func (r *repoPostgres) ListEntities(limit, offset uint64) ([]food.Food, error) {
	query,args,err := sq.Select("*").
		From(table).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil,err
	}
	rows,err := r.db.Query(query,args...)
	if err != nil {
		return nil,err
	}
	sliceFoods := []food.Food{}
	for rows.Next() {
		tmpFood := food.Food{}
		err = rows.Scan(&tmpFood.Id, &tmpFood.UserId, &tmpFood.Type, &tmpFood.Name, &tmpFood.PortionSize)
		if err != nil {
			return nil, err
		}
		sliceFoods = append(sliceFoods,tmpFood)
	}
	return sliceFoods,nil
}
func (r *repoPostgres) DescribeEntity(foodId uint64) (*food.Food, error) {
	query,args,err := sq.Select("id","user_id","type","name","portion_size").
		From(table).
		Where(sq.Eq{"id":foodId}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil,err
	}
	row := r.db.QueryRow(query,args...)
	f := food.Food{}
	err = row.Scan(&f.Id,&f.UserId,&f.Type,&f.Name,&f.PortionSize)

	if err != nil {
		return nil,err
	}
	return &f,nil
}
func (r *repoPostgres) RemoveEntity(foodId uint64) error {
	query,args,err := sq.Delete("*").
		From(table).
		Where(sq.Eq{"id":foodId}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	_,err = r.db.Exec(query,args...)
	if err != nil {
		return err
	}
	return nil
}