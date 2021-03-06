package topic

import (
	"database/db"
	"rutracker/topicData"
	"strconv"
	"strings"
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
	converted := []string{}
	for id := range data {
		name := strings.Split(topicData.GetName(id), "/")[0]
		converted = append(converted, strconv.Itoa(id) + ": " + name + "\n")
	}
	return converted
}

