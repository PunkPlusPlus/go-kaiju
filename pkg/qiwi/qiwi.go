package qiwi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"kaijuVpn/pkg/qiwi/models"
	"log"
	"net/http"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

const CREATE_BILL_URL = "https://api.qiwi.com/partner/bill/v1/bills/"

func CreateBill() models.CreateBillResponse {
	key := os.Getenv("QIWI_PRIVATE_KEY")

	request := createRequest(key)
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatalln(err.Error())
	}

	billResponse := parseReponse(response)

	return billResponse
}

func parseReponse(response *http.Response) models.CreateBillResponse {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	billResponse := models.CreateBillResponse{}
	err = json.Unmarshal(body, &billResponse)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return billResponse
}

func createRequest(key string) *http.Request {
	uuid := uuid.New().String()

	data := getData()
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest(http.MethodPut, CREATE_BILL_URL+uuid, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Authorization", "Bearer "+key)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	return request
}

func getData() models.CreateBillRequest {
	amount := models.Amount{}
	amount.Currency = "RUB"
	amount.Value = 10
	main := models.CreateBillRequest{}
	main.Amount = amount
	main.ExpirationDateTime = time.Now().Add(10 * time.Minute).Format("2006-01-02T15:04:05-0700")
	return main
}
func Pulling(billId string, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	key := os.Getenv("QIWI_PRIVATE_KEY")

	request, err := http.NewRequest(http.MethodGet, CREATE_BILL_URL+billId, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Authorization", "Bearer "+key)
	request.Header.Set("Accept", "application/json")
	for {
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		status_response := models.BillStatusResponse{}
		err = json.Unmarshal(body, &status_response)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(status_response.Status.Value)
		statuses := map[string]func(*tgbotapi.BotAPI, tgbotapi.Update){
			models.BILL_PAID:     BillStatusPaid,
			models.BILL_EXPIRED:  BillStatusExpired,
			models.BILL_REJECTED: BillStatusRejected,
		}

		if checker, ok := statuses[status_response.Status.Value]; ok {
			checker(bot, update)
			break
		}
		time.Sleep(2 * time.Second)
	}

}
func BillStatusPaid(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Спасибо за покупку!")
	bot.Send(msg)
}
func BillStatusExpired(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var expairedMessage = tgbotapi.NewMessage(
		update.CallbackQuery.Message.Chat.ID,
		"просрочен",
	)
	bot.Send(expairedMessage)
}
func BillStatusRejected(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Платёж отменён.")
	bot.Send(msg)
}
