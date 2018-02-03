package topic

import (
	"database/db"
)

func Add(userId int, topicId int) bool {
	userTopics, isset := db.GetData(userId)
	if isset == false {
		userTopics = []int{topicId}
	} else {
		userTopics = append(userTopics, topicId)
	}

	return db.SetData(userId, userTopics);
}

func GetList(userId int) []int {
	data, isset := db.GetData(userId)
	if !isset {
		return []int{}
	}
	return data
}

func Delete(userId int, topicId int) bool {
	data, isset := db.GetData(userId)
	if !isset {
		return true
	}

	newData := []int{}
	for _, value := range data {
		if value == topicId {
			continue
		}
		newData = append(newData, value)
	}

	return db.SetData(userId, newData)
}
