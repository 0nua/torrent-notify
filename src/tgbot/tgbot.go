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
const DELETE_COMMAND = "/delete"

var command = make(map[int]string)

type Dialog struct {
	ChatId  int64
	UserId  int
	Text    string
	Command string
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

		if dialog.Command != "" {
			topicId, err := strconv.Atoi(dialog.Text)
			error.Catch(err)
			message := "Подписал Вас на топик #" + dialog.Text
			switch dialog.Command {
			case TOPIC_COMMAND:
				topic.Add(dialog.UserId, topicId)

				break
			case DELETE_COMMAND:
				topic.Delete(dialog.UserId, topicId)
				message =  "Отписал Вас от топика #" + dialog.Text
				break
			}
			sendMessage(bot, message, dialog)
			delete(command, dialog.UserId)
			continue
		}

		if dialog.Text == TOPIC_COMMAND || dialog.Text == DELETE_COMMAND {
			sendMessage(bot, "Жду номер топика", dialog)
			command[dialog.UserId] = dialog.Text
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

	command, isset := command[dialog.UserId]
	if isset {
		dialog.Command = command
	} else {
		dialog.Command = ""
	}

	return *dialog
}

func sendMessage(bot tgbotapi.BotAPI, message string, dialog Dialog) {
	msg := tgbotapi.NewMessage(dialog.ChatId, message)
	bot.Send(msg)
}