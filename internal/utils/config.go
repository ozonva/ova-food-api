package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetConfigFromFiles(filepath string) {
	for i := 0; i < 5; i++ {
		func() {
			file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
			defer file.Close()

			if err != nil {
				panic(err)
			}
			conf, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			fmt.Printf("iteratoin:%d, config: %v\n", i, string(conf))
		}()
	}
}
