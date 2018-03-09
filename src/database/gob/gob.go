package gob

import (
	"os"
	"encoding/gob"
	"config"
	"error"
)

func Write(data map[int]map[int]int) bool {
	file, err := os.Create(config.GetDB())
	error.Catch(err)
	err = gob.NewEncoder(file).Encode(data)
	file.Close()

	return err == nil
}

func Read() map[int]map[int]int {
	data := make(map[int]map[int]int)
	file, err := os.Open(config.GetDB())
	if (err != nil) {
		return make(map[int]map[int]int)
	}
	err = gob.NewDecoder(file).Decode(&data)
	if err != nil {
		return make(map[int]map[int]int)
	}
 	file.Close()
	return data
}
