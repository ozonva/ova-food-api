package repo

import (
	"os"

	f "github.com/ozonva/ova-food-api/pkg/food"
)

type repoFile struct {
	filepath string
}

func (r *repoFile) UpdateFile(data string) error {
	file, err := os.OpenFile(r.filepath, os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(data)
	file.WriteString("\n")
	return nil
}
func (r *repoFile) AddEntities(entities []f.Food) error {
	var runes []rune
	for _, en := range entities {
		runes = append(runes, []rune(en.String())...)
	}
	err := r.UpdateFile(string(runes))
	if err != nil {
		return err
	}
	return nil
}
func (r *repoFile) ListEntities(limit, offset uint64) ([]f.Food, error) {
	return nil, nil
}
func (r *repoFile) DescribeEntity(entityId uint64) (*f.Food, error) {
	return nil, nil
}
