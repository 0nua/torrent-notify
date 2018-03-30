package checker

import (
	"database/db"
	"rutracker/topicData"
	"gopkg.in/telegram-bot-api.v4"
	"time"
	"rutracker/topic"
)

func Run(bot tgbotapi.BotAPI) {
	for {
		data := db.Get()
		for userId, topics := range data {
			for topicId, size := range topics {
				newSize := topicData.GetSize(topicId, true)
				if newSize > size {
					msg := tgbotapi.NewMessage(int64(userId), topicData.GetName(topicId) + " обновлён!")
					bot.Send(msg)

					topic.Add(userId, topicId)
				}
				time.Sleep(5*time.Second)
			}
		}
	}
}
