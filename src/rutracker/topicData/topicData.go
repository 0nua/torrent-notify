package topicData

import (
	"rutracker/rutrackerClient"
	"strconv"
	"fmt"
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
	return  getData(topicId).Topic_title
}

func getData(topicId int) TopicData {
	if data, isset := topicData[topicId]; isset {
		return data
	}

	topicData[topicId] = getTopicData(topicId)
	return topicData[topicId]
}

func getTopicData(topicId int) TopicData {
	fmt.Println("Fetchig json")
	result := rutrackerClient.Request(METHOD_GET_TOPIC_DATA, topicId)
	data := result[strconv.Itoa(topicId)].(map[string]interface{})
	return convertData(data)
}

func convertData(data map[string]interface{}) TopicData {
	return TopicData{
		Size: int(data["size"].(float64)),
		Topic_title: data["topic_title"].(string),
	}
}