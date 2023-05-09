package handlers

import (
	"kaijuVpn/pkg/database/users"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(command string) (func(*tgbotapi.BotAPI, tgbotapi.Update), error) {
	commands := map[string]func(*tgbotapi.BotAPI, tgbotapi.Update){
		"hello":    handleHello,
		"set_name": setName,
		"start":    startHandler,
	}
	if checker, ok := commands[command]; ok {
		return checker, nil
	}

	return defaultHandler, nil
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
	var telegram_user = update.Message.From
	var user = users.User{
		Telegram_id: strconv.FormatInt(telegram_user.ID, 10),
	}
	users.CreateIfNotExist(user)

	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Оплатить💰", "agree_purscase"),
			tgbotapi.NewInlineKeyboardButtonData("Как установить vpn💡", "noCallback"),
		),
	)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я - бот для продажи доступа к Kaiju VPN. Для оплаты нажми кнопку \"Оплатить\" и следуй инструкциям.")
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)
}
