package food

import (
	"encoding/json"
)

const (
	Unknown = iota
	Drinks
	Foods
)

type Food struct {
	Id          uint64
	UserId      uint64
	Type        uint8
	Name        string
	PortionSize float32
}

func (f Food) String() string {
	mf, err := json.Marshal(f)
	if err != nil {
		panic(err.Error())
	}
	return string(mf)
}

func CreateFood(foodInfo []byte) *Food {
	var food Food
	err := json.Unmarshal(foodInfo, &food)
	if err != nil {
		panic(err.Error())
	}
	return &food
}
