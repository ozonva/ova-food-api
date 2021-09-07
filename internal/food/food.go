package food

import (
	"encoding/json"
	"unsafe"
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

func (f Food) size() int {
	size := 2 * 8       //2x uint64
	size += len(f.Name) //string
	size += 1           //uint8
	size += 4           //float32
	return size
}

func SizeFoods(f []Food) int {
	size := 0
	f = f[:cap(f)]
	size += cap(f) * int(unsafe.Sizeof(f))
	for i := range f {
		size += (f[i]).size()
	}
	return size
}
