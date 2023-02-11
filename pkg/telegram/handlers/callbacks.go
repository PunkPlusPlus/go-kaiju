package handlers

import (
	"errors"
	"kaijuVpn/pkg/qiwi"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCallback(callback string) (func(*tgbotapi.BotAPI, tgbotapi.Update), error) {

	if callback == "" {
		return nil, errors.New("Empty command")
	}

	switch callback {
	case "agree_purscase":
		return handlePurscase, nil
	default:
		return defaultHandler, nil
	}
}

func handlePurscase(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	result := qiwi.CreateBill()
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Ваша ссылка на оплату: \n"+result.PayUrl)

	bot.Send(msg)
}
