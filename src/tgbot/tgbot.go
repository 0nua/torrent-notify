package tgbot

import (
	"fail"
	"gopkg.in/telegram-bot-api.v4"
	"config"
	"strconv"
	"strings"
	"rutracker/topic"
	"rutracker/checker"
)

const TOPIC_COMMAND = "topic"
const GET_COMMAND = "list"
const DELETE_COMMAND = "delete"
const START_COMMAND = "start"
const CANCEL_COMMAND = "cancel"

const KEYBOARD_CANCEL = "cancel"
const KEYBOARD_MENU   = "menu"

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

	go checker.Run(bot)

	for update := range updatesChain {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		dialog, exist := assembleUpdate(update)
		if !exist {
			continue
		}

		if botCommand := getCommand(update); botCommand != "" {
			doBotCommand(botCommand, bot, dialog)
			continue
		}

		processDialogInput(bot, dialog)
	}
}

func assembleUpdate(update tgbotapi.Update) (Dialog, bool) {
	dialog := new(Dialog)

	if (update.Message != nil) {
		dialog.ChatId = update.Message.Chat.ID
		dialog.UserId = int(update.Message.Chat.ID)
		dialog.Text = update.Message.Text
	} else if (update.CallbackQuery != nil) {
		dialog.ChatId = update.CallbackQuery.Message.Chat.ID
		dialog.UserId = int(update.CallbackQuery.Message.Chat.ID)
		dialog.Text = ""
	} else {
		return *dialog, false
	}

	command, isset := command[dialog.UserId]
	if isset {
		dialog.Command = command
	} else {
		dialog.Command = ""
	}

	return *dialog, true
}

func sendMessage(bot tgbotapi.BotAPI, message string, dialog Dialog, keyboard string) {
	msg := tgbotapi.NewMessage(dialog.ChatId, message)
	if keyboard != "" {
		addKeyboard(&msg, keyboard)
	}
	bot.Send(msg)
}

func getCommand(update tgbotapi.Update) string {
	if (update.Message != nil) {
		if (update.Message.IsCommand()) {
			return update.Message.Command()
		}
	} else if (update.CallbackQuery != nil) {
		return update.CallbackQuery.Data
	}

	return ""
}

func addKeyboard(msg *tgbotapi.MessageConfig, keyboard string) {
	switch keyboard {
	case KEYBOARD_MENU:
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Список", GET_COMMAND),
				tgbotapi.NewInlineKeyboardButtonData("+", TOPIC_COMMAND),
				tgbotapi.NewInlineKeyboardButtonData("-", DELETE_COMMAND),
			),
		)
	case KEYBOARD_CANCEL:
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Отменить", CANCEL_COMMAND),
			),
		)

	}
}

func doBotCommand(botCommand string, bot tgbotapi.BotAPI, dialog Dialog) {
	switch botCommand {
	case TOPIC_COMMAND, DELETE_COMMAND:
		sendMessage(bot, "Жду номер топика для " + botCommand, dialog, KEYBOARD_CANCEL)
		command[dialog.UserId] = botCommand
	case GET_COMMAND:
		list := topic.GetList(dialog.UserId)
		message := "Список пуст"
		if len(list) != 0 {
			message = strings.Trim(strings.Join(list, ""), "[]")
		}

		sendMessage(bot, message, dialog, KEYBOARD_MENU)
	case START_COMMAND:
		msg := tgbotapi.NewMessage(dialog.ChatId, "Добрый день, я готов уведомлять Вас о новых сериях сериалов")
		addKeyboard(&msg, KEYBOARD_MENU)
		bot.Send(msg)
	case CANCEL_COMMAND:
		sendMessage(bot, "Что-то пошло не так?! Давай ещё разок", dialog, KEYBOARD_MENU)
	}
}

func processDialogInput(bot tgbotapi.BotAPI, dialog Dialog) {
	if dialog.Command == "" {
		return
	}

	message := "Подписал Вас на топик #" + dialog.Text
	switch dialog.Command {
	case TOPIC_COMMAND:
		topicId, err := strconv.Atoi(dialog.Text)
		fail.Catch(err)
		name := topic.Add(dialog.UserId, topicId)
		if (name != "") {
			message = message + " (" + name + ")"
		} else {
			message = "Такого топика не существует!"
		}
	case DELETE_COMMAND:
		topicId, err := strconv.Atoi(dialog.Text)
		fail.Catch(err)
		topic.Delete(dialog.UserId, topicId)
		message = "Отписал Вас от топика #" + dialog.Text
	}

	sendMessage(bot, message, dialog, KEYBOARD_MENU)
	delete(command, dialog.UserId)
}