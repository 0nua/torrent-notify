package tgbot

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	"config"
	"db"
	"strconv"
)

func Init() {
	waitForNumber := false
	bot, err := tgbotapi.NewBotAPI(config.GetToken())
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updatesChain, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Panic(err)
	}

	for update := range updatesChain {
		if update.Message == nil {
			continue
		}

		text := update.Message.Text;

		if waitForNumber {
			topicId, err := strconv.Atoi(update.Message.Text)
			if err != nil {
				log.Panic(err)
			}
			db.AddTopic(update.Message.From.ID, topicId)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Подписал Вас на топик #" + update.Message.Text)
			bot.Send(msg)
		}


		if text == "/subscribe" {
			waitForNumber = true;
			msg:= tgbotapi.NewMessage(update.Message.Chat.ID, "Жду номер топика!")
			bot.Send(msg)
		}
	}
}