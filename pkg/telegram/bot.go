package telegram

import (
	"kaijuVpn/pkg/telegram/handlers"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Run() error {
	err := b.createWebHook(os.Getenv("WEBHOOK_URL")+"/", b.bot.Token)
	if err != nil {
		log.Fatal(err)
	}
	info, err := b.bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	b.handleWebHookError(info)

	updates := b.bot.ListenForWebhook("/" + b.bot.Token)
	go http.ListenAndServe("0.0.0.0:4000", nil)

	b.handleUpdates(updates)
	return err
}

func (b *Bot) handleWebHookError(info tgbotapi.WebhookInfo) {
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
}

func (b *Bot) createWebHook(url string, token string) error {
	wh, err := tgbotapi.NewWebhook(url + token)

	_, err = b.bot.Request(wh)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.CallbackQuery != nil {
			handler, err := handlers.HandleCallback(update.CallbackQuery.Data)
			if err == nil {
				go handler(b.bot, update)
			}
		} else if update.Message != nil {
			command := update.Message.Command()
			handler, err := handlers.HandleCommand(command)
			if err == nil {
				go handler(b.bot, update)
			}
		}
	}
}
