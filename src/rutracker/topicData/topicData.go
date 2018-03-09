package topicData

import (
	"rutracker/rutrackerClient"
	"strconv"
)

const METHOD_GET_TOPIC_DATA = "get_tor_topic_data"

var topicData = map[int]TopicData{}

type TopicData struct {
	Size int
	Topic_title string
}

func GetSize(topicId int) int {
	return getData(topicId).Size
}

func GetName(topicId int) string {
	return getData(topicId).Topic_title
}

func getData(topicId int) TopicData {
	if data, exist := topicData[topicId]; exist {
		if (data.Size != 0) {
			return data
		}
	}

	return getTopicData(topicId);
}

func getTopicData(topicId int) TopicData {
	result := rutrackerClient.Request(METHOD_GET_TOPIC_DATA, topicId)

	key := strconv.Itoa(topicId)
	if (result[key] == nil) {
		return TopicData{}
	}
	data := result[key].(map[string]interface{})
	return convertData(data)
}

func convertData(data map[string]interface{}) TopicData {
	return TopicData{
		Size: int(data["size"].(float64)),
		Topic_title: data["topic_title"].(string),
	}
}