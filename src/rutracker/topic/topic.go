package topic

import (
	"database/db"
	"rutracker/topicData"
)

func Add(userId int, topicId int) string {
	userTopics, isset := db.GetData(userId)
	if isset == false {
		userTopics = map[int]int{}
	}

	size := topicData.GetSize(topicId)
	if size != 0 {
		userTopics[topicId] = size
		db.SetData(userId, userTopics);
	}

	return topicData.GetName(topicId)
}

func GetList(userId int) []int {
	data, isset := db.GetData(userId)
	if !isset {
		return []int{}
	}
	return convert(data)
}

func Delete(userId int, topicId int) bool {
	data, isset := db.GetData(userId)
	if !isset {
		return true
	}

	newData := map[int]int{}
	for id, updated := range data {
		if id == topicId {
			continue
		}
		newData[id] = updated
	}

	return db.SetData(userId, newData)
}

func convert(data map[int]int) []int {
	conveted := []int{}
	for id := range data {
		conveted = append(conveted, id)
	}
	return conveted
}

