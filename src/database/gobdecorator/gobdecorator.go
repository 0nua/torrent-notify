package gobdecorator

import (
	"database/gob"
)

func Save(data map[int][]int) bool {
	return gob.Write(data)
}

func Get() map[int][]int {
	return gob.Read()
}

func merge(savedData map[int][]int, newData map[int][]int) map[int][]int {
	for userId, topics := range newData {
		savedTopics, isset := savedData[userId]
		if (isset == true) {
			savedData[userId] = append(savedTopics, topics...)
		} else {
			savedData[userId] = topics
		}
	}

	return savedData
}