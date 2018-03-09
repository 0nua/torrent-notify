package topicData

import (
	"rutracker/rutrackerClient"
)

var data map[string]interface{} = nil

func GetSize(topicId int) int {
	topic := getData(topicId)
	return int(topic["size"].(float64))
}

func GetName(topicId int) string {
	topic := getData(topicId)
	return topic["topic_title"].(string)
}

func getData(topicId int) map[string]interface{} {
	if (data == nil) {
		data = rutrackerClient.GetTopicData(topicId)
	}

	return data
}
