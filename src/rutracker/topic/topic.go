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
