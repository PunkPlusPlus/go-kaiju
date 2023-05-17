package handlers

import (
	"kaijuVpn/pkg/qiwi"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var oldMarkup tgbotapi.InlineKeyboardMarkup

func HandleCallback(callback string) (func(*tgbotapi.BotAPI, tgbotapi.Update), error) {
	callbacks := map[string]func(*tgbotapi.BotAPI, tgbotapi.Update){
		"agree_purscase": handlePurscase,
		"back":           back,
	}
	if checker, ok := callbacks[callback]; ok {
		return checker, nil
	}

	return nil, nil
}

func handlePurscase(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var waitMessage = tgbotapi.NewEditMessageText(
		update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID,
		"Ожидайте ссылку для оплаты ⏳",
	)
	bot.Send(waitMessage)
	result := qiwi.CreateBill()
	var buttons = []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonURL("Ссылка для оплаты💸", result.PayUrl),
		tgbotapi.NewInlineKeyboardButtonData("Назад⮑", "back"),
	}
	oldMarkup = *update.CallbackQuery.Message.ReplyMarkup
	var keyboard = tgbotapi.NewEditMessageReplyMarkup(
		update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID,
		tgbotapi.NewInlineKeyboardMarkup(buttons),
	)

	bot.Send(keyboard)

	go qiwi.Pulling(result.BillId, bot, update)
}

func back(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var msg = tgbotapi.NewEditMessageReplyMarkup(
		update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID,
		oldMarkup,
	)
	bot.Send(msg)
}
