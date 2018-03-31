package checker

import (
	"database/db"
	"rutracker/topicData"
	"gopkg.in/telegram-bot-api.v4"
	"time"
	"rutracker/topic"
	"config"
)

func Run(bot tgbotapi.BotAPI) {
	for {
		data := db.Get()
		for userId, topics := range data {
			for topicId, size := range topics {
				newSize := topicData.GetSize(topicId, true)
				if newSize > 0 && newSize > size {
					msg := tgbotapi.NewMessage(int64(userId), getMessage(topicId))
					bot.Send(msg);

					topic.Add(userId, topicId)
				}
				sleep()
			}
		}
	}
}

func getMessage(topicId int) string {
	return topicData.GetName(topicId) + " обновлён!";
}

func sleep() {
	timeout := config.GetUpdateTimeout();
	time.Sleep(time.Duration(timeout) * time.Second)
}
