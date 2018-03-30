package topic

import (
	"database/db"
	"rutracker/topicData"
	"strconv"
)

func Add(userId int, topicId int) string {
	userTopics, isset := db.GetData(userId)
	if isset == false {
		userTopics = map[int]int{}
	}

	size := topicData.GetSize(topicId, false)
	if size != 0 {
		userTopics[topicId] = size
		db.SetData(userId, userTopics);
	}

	return topicData.GetName(topicId)
}

func GetList(userId int) []string {
	data, isset := db.GetData(userId)
	if !isset {
		return []string{}
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

func convert(data map[int]int) []string {
	conveted := []string{}
	for id := range data {
		name := topicData.GetName(id)
		conveted = append(conveted, strconv.Itoa(id) + ": " + name + "\n")
	}
	return conveted
}

