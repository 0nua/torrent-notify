package rutrackerClient

import (
	"config"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fail"
	"strconv"
)

const PARAM_BY = "by"
const PARAM_VAL = "val"

type jsonObject struct {
	Result map[string]interface{} `json:"result"`
}

func Request(method string, topicId int) map[string]interface{} {
	url := getUrl(method, prepareParams(strconv.Itoa(topicId)))

	response, err := http.Get(url)
	fail.Catch(err)

	body, readErr := ioutil.ReadAll(response.Body)
	fail.Catch(readErr)
	response.Body.Close()

	jsonObject := jsonObject{}
	jsonErr := json.Unmarshal(body, &jsonObject)
	fail.Catch(jsonErr)

	return jsonObject.Result
}

func getUrl(method string, params map[string]string) string {
	parts := []string{
		config.GetRutrackerApi(),
		method,
	}

	return strings.Join(parts, "/") + parseParams(params)
}

func parseParams(params map[string]string) string {
	items := []string{}
	for name, value := range params {
		items = append(items, name + "=" + value)
	}

	return "?" + strings.Join(items, "&")
}

func prepareParams(topicId string) map[string]string {
	return map[string]string{
		PARAM_BY: "topic_id",
		PARAM_VAL: topicId,
	}
}
