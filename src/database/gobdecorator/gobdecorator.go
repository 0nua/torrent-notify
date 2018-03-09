package gobdecorator

import (
	"database/gob"
)

func Save(data map[int]map[int]int) bool {
	return gob.Write(data)
}

func Get() map[int]map[int]int {
	return gob.Read()
}