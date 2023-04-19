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
	main.ExpirationDateTime = time.Now().Add(10*time.Minute).Format("2006-01-02T15:04:05") + "+00:00"
	return main
}