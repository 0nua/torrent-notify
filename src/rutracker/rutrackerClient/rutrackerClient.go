package rutrackerClient

import (
	"config"
	"strings"
	"strconv"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

const METHOD_GET_TOPIC_DATA = "get_tor_topic_data"

func GetTopicData(topicId int) map[string]interface{} {
	key := strconv.Itoa(topicId)
	result := request(METHOD_GET_TOPIC_DATA, prepareParams(key))
	return result[key].(map[string]interface{})
}

func getUrl(method string, params map[string]string) string {
	parts := []string{
		config.GetRutrackerApi(),
		method,
	}

	return strings.Join(parts, "/") + parseParams(params)
}

func parseParams(params map[string]string) string {
	tmpParams := []string{}
	for name, value := range params {
		tmpParams = append(tmpParams, name + "=" + value)
	}

	return "?" + strings.Join(tmpParams, "&")
}

func prepareParams(topicId string) map[string]string {
	params := map[string]string{}
	params["by"] = "topic_id"
	params["val"] = topicId

	return params
}

func request(method string, params map[string]string) map[string]interface{} {
	url := getUrl(method, params)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if (readErr != nil) {
		log.Fatal(readErr)
	}

	jsonObject := map[string]interface{}{}
	jsonErr := json.Unmarshal(body, &jsonObject)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	response.Body.Close()

	return jsonObject["result"].(map[string]interface{})
}