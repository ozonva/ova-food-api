package utils

import (
	"fmt"
	"os"
)

func UpdateConfig(filepath string) {
	updateConfig := func(iteration int, filepath string) error {
		file, err := os.OpenFile(filepath, os.O_APPEND, 0755)
		if err != nil {
			return err
		}
		defer file.Close()
		config := fmt.Sprintf("config %d = %d\n", iteration, iteration*2)
		file.WriteString(config)
		return nil
	}

	for i := 0; i < 5; i++ {
		err := updateConfig(i, filepath)
		if err != nil {
			panic(err)
		}
	}
}
