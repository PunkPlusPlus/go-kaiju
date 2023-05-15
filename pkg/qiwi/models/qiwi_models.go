package models

const (
	BILL_WAITING  = "WAITING"
	BILL_PAID     = "PAID"
	BILL_REJECTED = "REJECTED"
	BILL_EXPIRED  = "EXPIRED"
)

type CreateBillResponse struct {
	SiteId             string     `json:"siteId"`
	BillId             string     `json:"billId"`
	Amount             Amount     `json:"amount"`
	Status             BillStatus `json:"status"`
	Customer           Customer   `json:"customer"`
	CreationDateTime   string     `json:"creationDateTime"`
	ExpirationDateTime string     `json:"expirationDateTime"`
	PayUrl             string     `json:"payUrl"`
}

type BillStatus struct {
	Value           string `json:"value"`
	ChangedDateTime string `json:"changedDateTime"`
}
type BillStatusResponse struct {
	Status BillStatus `json:"status"`
}

type Customer struct {
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Account string `json:"account"`
}

// ГГГГ-ММ-ДДTчч:мм:сс±чч:мм
type CreateBillRequest struct {
	Amount             Amount `json:"amount"`
	ExpirationDateTime string `json:"expirationDateTime"`
}

type Amount struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value,string"`
}
