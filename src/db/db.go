package db

import (
	"error"
	"io/ioutil"
	"encoding/json"
	"config"
)

type User struct {
	UserId int `json:"id"`
	Topics []int `json:"topics"`
}

type DB struct {
	Items map[int]User
}

func AddTopic(userId int, topicId int) bool {
	var topics []int
	db := getDB()
	user, isset := db.Items[userId]
	if (isset) {
		topics = user.Topics
	}

	user.Topics = append(topics, topicId)
	db.Items[userId] = user

	return saveDB(db)
}

func getDB() DB {
	jsonString, err := ioutil.ReadFile(config.GetDB())
	error.Check(err)

	db := DB{}
	json.Unmarshal(jsonString, db)
	return db;
}

func saveDB(db DB) bool {
	jsonString, err := json.Marshal(db)
	error.Check(err)

	return ioutil.WriteFile(config.GetDB(), jsonString, 755) == nil
}
