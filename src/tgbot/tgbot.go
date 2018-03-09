package tgbot

import (
	"fail"
	"gopkg.in/telegram-bot-api.v4"
	"config"
	"strconv"
	"fmt"
	"strings"
	"rutracker/topic"
)

const TOPIC_COMMAND = "topic"
const GET_COMMAND = "list"
const DELETE_COMMAND = "delete"

var command = make(map[int]string)

type Dialog struct {
	ChatId  int64
	UserId  int
	Text    string
	Command string
}

func Init() {
	bot, err := tgbotapi.NewBotAPI(config.GetToken())
	fail.Catch(err)
	topicSaver(*bot)
}

func topicSaver(bot tgbotapi.BotAPI) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updatesChain, err := bot.GetUpdatesChan(updateConfig)
	fail.Catch(err)

	for update := range updatesChain {
		if update.Message == nil {
			continue
		}

		dialog := assembleUpdate(update)

		if dialog.Command != "" {
			topicId, err := strconv.Atoi(dialog.Text)
			fail.Catch(err)
			message := "Подписал Вас на топик #" + dialog.Text
			switch dialog.Command {
			case TOPIC_COMMAND:
				name := topic.Add(dialog.UserId, topicId)
				if (name != "") {
					message = message + " (" + name + ")"
				} else {
					message = "Такого топика не существует!"
				}
				break
			case DELETE_COMMAND:
				topic.Delete(dialog.UserId, topicId)
				message = "Отписал Вас от топика #" + dialog.Text
				break
			}
			sendMessage(bot, message, dialog)
			delete(command, dialog.UserId)
			continue
		}


		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case TOPIC_COMMAND, DELETE_COMMAND:
				sendMessage(bot, "Жду номер топика", dialog)
				command[dialog.UserId] = update.Message.Command()
				break
			case GET_COMMAND:
				message := strings.Trim(
					strings.Join(
						strings.Split(
							fmt.Sprint(
								topic.GetList(dialog.UserId)), " "), ","), "[]")
				sendMessage(bot, message, dialog)
				break
			}
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