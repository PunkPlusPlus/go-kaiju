package handlers

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(command string) (func(*tgbotapi.BotAPI, tgbotapi.Update), error) {

	if command == "" {
		return nil, errors.New("Empty command")
	}

	switch command {
	case "hello":
		return handleHello, nil
	case "set_name":
		return setName, nil
	case "start":
		return startHandler, nil
	default:
		return defaultHandler, nil
	}
}

func handleHello(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello Get!")
	bot.Send(msg)
}

func setName(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "SET_NAME COMMAND")
	bot.Send(msg)
}

func defaultHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "DEFAULT HANDLER")
	bot.Send(msg)
}

func startHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Yes", "agree_purscase"),
			tgbotapi.NewInlineKeyboardButtonData("No", "noCallback"),
		),
	)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)
}
