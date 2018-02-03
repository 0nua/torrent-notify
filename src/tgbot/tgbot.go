package tgbot

import (
	"error"
	"gopkg.in/telegram-bot-api.v4"
	"config"
	"strconv"
	"fmt"
	"strings"
	"rutracker/topic"
)

const TOPIC_COMMAND = "/topic"
const GET_COMMAND = "/list"

var subscribe = make(map[int]bool)

type Dialog struct {
	ChatId int64
	UserId int
	Text   string
	IsWait bool
}

func Init() {
	bot, err := tgbotapi.NewBotAPI(config.GetToken())
	error.Catch(err)
	topicSaver(*bot)
}

func topicSaver(bot tgbotapi.BotAPI) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updatesChain, err := bot.GetUpdatesChan(updateConfig)
	error.Catch(err)

	for update := range updatesChain {
		if update.Message == nil {
			continue
		}

		dialog := assembleUpdate(update)

		if dialog.IsWait == true {
			topicId, err := strconv.Atoi(dialog.Text)
			error.Catch(err)
			topic.Add(dialog.UserId, topicId)
			sendMessage(bot, "Подписал Вас на топик #" + dialog.Text, dialog)
			subscribe[dialog.UserId] = false
			continue
		}

		if dialog.Text == TOPIC_COMMAND {
			sendMessage(bot, "Жду номер топика", dialog)
			subscribe[dialog.UserId] = true
			continue
		}

		if dialog.Text == GET_COMMAND {
			message := strings.Trim(
				strings.Join(
					strings.Split(
						fmt.Sprint(
							topic.GetList(dialog.UserId)), " "), ","), "[]")
			sendMessage(bot, message, dialog)
			continue
		}
	}
}

func assembleUpdate(update tgbotapi.Update) Dialog {
	dialog := new(Dialog)

	dialog.ChatId = update.Message.Chat.ID
	dialog.UserId = update.Message.From.ID
	dialog.Text = update.Message.Text
	dialog.IsWait = subscribe[dialog.UserId]

	return *dialog
}

func sendMessage(bot tgbotapi.BotAPI, message string, dialog Dialog) {
	msg := tgbotapi.NewMessage(dialog.ChatId, message)
	bot.Send(msg)
}